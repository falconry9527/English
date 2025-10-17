## EVM的数据存储
```
Memory：函数调用期间的临时存储，存放函数参数、局部变量，中间计算结果，调用结束后清空。
Stack：EVM 执行计算的核心区域，存储值类型的 局部变量，中间计算结果，生命周期极短，最多 1024 个 slot。
Storage：合约的永久链上状态存储，存放全局变量、映射（mapping）、可变数组、结构体等，操作成本高，需要消耗 Gas。
Calldata：外部函数调用的 只读输入参数存储区，存放 address、uint、bytes、string 等，生命周期仅在函数执行期间，成本低且不可修改。

全局变量：默认在 storage（包括 mapping、array、struct、值类型）
局部变量：
值类型 → 默认 stack
引用类型 → 必须指定 memory 或 storage（数组/struct）

mapping →  只能是 合约级全局变量（contract-level state),只能存储在 stroage,不能在函数中定义
栈上存储的是 mapping 起始 slot 索引，访问 key 时通过 keccak256(key, slot) 定位真实 storage

数组/struct storage 引用：栈上存储的是起始  slot 索引，访问元素时根据 slot + 偏移计算 storage

```

## CREATE2
```
定义: 它们都是 EVM 中的合约创建指令（opcode）。

一. CREATE : 合约地址是不可预测的，地址由 部署者地址 + nonce 决定 
主要 依赖于部署顺序和次数（nonce） 
address = keccak256(rlp(sender_address, nonce))[12:]
sender_address：部署者地址
nonce：部署者账户的交易计数（每次交易 +1）

二. CREATE2 : 在部署前就能预测合约地址，地址由部署者地址 + 合约字节码 + 盐值salt 决定
主要依赖于 盐值(salt: 手动传入的参数) 。
解决 主网与测试网部署地址不一致的问题；

计算方式:
address = keccak256(
    0xFF,
    sender,
    salt,
    keccak256(bytecode)
)[12:]  // 取后 20 字节
取决于 ： 部署者地址 + 盐值（salt） + 合约字节码

部署者地址（sender） → 部署者地址（比如 Factory 合约地址）
盐值（salt） → 由部署者自定义（常用 keccak256(token0, token1)）
合约字节码（bytecode) → 待部署合约的字节码（creationCode）

deploy 使用 create2 真正部署(assembly 调用 CREATE2)；
getAddress 用相同参数计算地址；

三. 案例:
uniswap 的交易对池子是单独的合约，是唯一的，固定的。
使用 CREATE2 创建的，

四. new 
使用 new 创建合约的时候，
不传 盐值（salt）,底层调用的就是 CREATE；
传入 盐值（salt）,底层调用的就是 CREATE2 
例如: new Pool{salt: salt}() 
```

## abi.encode 和 abi.encodePacked 之间有什么区别？
```
bytes memory data = abi.encode(a, b, c);
将输入参数按 ABI 编码（Solidity 标准 ABI）序列化，返回类型：bytes memory
不可逆：每个参数都有固定大小或长度前缀
案例：函数调用、签名消息

不可逆：每个参数紧凑打包，没有固定大小或长度前缀
解码时可能会出现二义性（例如两个动态类型参数拼在一起）
案例 ： 生成唯一标识
uniswap 的交易对池子是单独的合约，是唯一的，固定的 。
使用 encodePacked，uniswap中,生成交易对的唯一标识-盐值
```

## Solidity 函数选择器
```
一. 定义
函数选择器是 Solidity 函数的唯一标识 。
它是 函数签名 functionName(type1,type2,...) 的 Keccak-256 哈希 的前 4 个字节。
bytes4 selector = bytes4(keccak256("transfer(address,uint256)"));

主要用于底层函数 call 和 delegatecall 的调用

4字节存储（232≈4.3×109 43 亿个值）

二.管理策略
函数数量 ≤4：顺序 if-else 比较，O(n) 时间复杂度。
函数数量 >4：编译器自动生成 switch-case 跳转表，直接匹配选择器，O(1) 调度。
函数数量很多（几十个以上）：采用 分段/二分查找逻辑，时间复杂度 O(log n)，保证调度高效。

如果选择器不匹配任何函数，跳转到 fallback 或 revert，保证安全。
```

## EIP-1559
```
1559是以太坊的一个升级，改变 gas 计算规则
之前：以太坊交易的成本由矿工通过拍卖机制来决定
之后：交易的成本由两个因素决定：Gas Price和Gas Limit
```

## 荷兰式拍卖
```
荷兰式拍卖是 价格从高到低下降直到有人接受，而英式拍卖是 价格从低开始逐渐被竞价者抬高直到无人加价。
```

## tx.origin
```
msg.sender → 调用当前合约的直接调用者
tx.origin → 发起整个交易的最初外部账户

为什么不用 tx.origin 作为鉴权：
防止钓鱼攻击：
攻击者部署中间，诱导原始用户A，调用 转账功能。
```

## 三明治交易
```
三明治攻击是攻击者监听 mempool（未打包交易池）中即将执行的大额交易，
在 大概买入单 前后 夹入 买单 和 卖单 ，进行套利

mempool：当用户发送交易时，交易 首先到达节点的内存池（mempool），交易在这里等待矿工或验证者打包进区块；

防御：拆分大单交易，避免滑点过大
```

## 冷读（cold read）和热读（warm read）
```
冷读是指第一次读取存储变量时，需要从存储中读取变量的值，这需要较高的gas费用。
热读是指再次读取存储变量时，可以从缓存中读取变量的值，这需要较低的gas费用。
热读和冷读是由Ethereum虚拟机（EVM）自动处理的。
```

## modifier
```
modifier是一种函数修饰符，用于声明一个 函数修改器,比如 onlyowner
```

## fallback 和 receive 的区别？
```
receive()：接收 ETH ；
fallback()：处理未匹配函数或 ETH 调用。
```


## payable
```
payable函数是一种特殊类型的函数，允许合约接受以太币作为支付
payable(address)
```




