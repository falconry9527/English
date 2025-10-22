## Data storage for EVM
```
----- structure of evn
Memory: A temporary storage area used during function execution. 
It stores function parameters, local variables, and intermediate computation results.

Stack: The computation area of the EVM.
It stores local variables and intermediate computation results. 
It has a very short lifetime and can hold up to 1024 slots.

Storage: Permanent on-chain storage.
It stores global variables such as mappings and dynamic arrays.

Calldata: Read-only data.
It stores function parameters for external functions.

----- Global and local Variables
Global Variables : stored in the storage by default.
local Variables :
Value types: stored in the stack by default.(int, bool)
Reference types : Must explicitly specify memory or storage. (arrays, structs)

----- Mapping 
Mapping: it can only be a contract-level global variable, stored in storage, 
and cannot be defined within a function.
1. The stack stores a starting index of a slot, which points to a slot in storage.
2. Index of each element: keccak256(key, slot)

```

## CREATE2
```
定义: 它们都是 EVM 中的合约创建指令（opcode）。

一. CREATE : 合约地址是不可预测的，地址由 部署者地址 + nonce 决定 
主要 依赖于部署顺序和次数（nonce） 
address = keccak256(rlp(sender_address, nonce))[12:]
sender_address:部署者地址
nonce:部署者账户的交易计数（每次交易 +1）

二. CREATE2 : 在部署前就能预测合约地址，解决 主网与测试网部署地址不一致的问题
地址由部署者地址 + 合约字节码 + 盐值salt 决定
主要依赖于 盐值(salt: 手动传入的参数) 。
解决 主网与测试网部署地址不一致的问题；

计算方式:
address = keccak256(
    0xFF,
    sender,
    salt,
    keccak256(bytecode)
)[12:]  // 取后 20 字节
取决于 : 部署者地址 + 盐值（salt） + 合约字节码

部署者地址（sender） : 部署者地址（比如 Factory 合约地址）
盐值（salt） : 由部署者自定义（常用 keccak256(token0, token1)）
合约字节码（bytecode) : 待部署合约的字节码（creationCode）

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
将输入参数按 ABI 编码（Solidity 标准 ABI）序列化，返回类型:bytes memory
不可逆:每个参数都有固定大小（32字节填充，刚好放在一个slot）或长度前缀
案例:函数调用、签名消息

不可逆:每个参数紧凑打包，没有固定大小（没有做32字节填充），长度前缀
解码时可能会出现二义性（例如两个动态类型参数拼在一起）
案例 : 生成唯一标识
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
函数数量 ≤4:顺序 if-else 比较，O(n) 时间复杂度。
函数数量 >4:编译器自动生成 switch-case 跳转表，直接匹配选择器，O(1) 调度。
函数数量很多（几十个以上）:采用 分段/二分查找逻辑，时间复杂度 O(log n)，保证调度高效。

如果选择器不匹配任何函数，跳转到 fallback 或 revert，保证安全。
```

## delegatecall 和 call 的不同
```
不同一:
call 在被调用者上下文执行，修改被调用者状态。
例如:合约A调用合约B的函数，函数是在 B的上下文执行，修改b的变量

delegatecall 在调用者上下文执行，修改调用者状态。
例如: 可升级合约的代理合约，转发用户的函数给实现合约，函数是在 代理合约 的上下文执行，修改 代理合约 的变量 ；

不同二:
delegatecall: msg.sender 是实际调用者
call: msg.sender 是 直接调用者

用户A->合约B--delegatecall/call--> 合约C
合约C 的  msg.sender
delegatecall : 用户A
call : 合约B
```

## EIP-1559
```
1559是以太坊的一个升级，改变 gas 计算规则
之前:以太坊交易的成本由矿工通过拍卖机制来决定
之后:交易的成本由两个因素决定:Gas Price和Gas Limit
```

## 荷兰式拍卖
```
荷兰式拍卖是 价格从高到低下降直到有人接受，而英式拍卖是 价格从低开始逐渐被竞价者抬高直到无人加价。
```

## tx.origin
```
msg.sender : 调用当前合约的直接调用者
tx.origin : 发起整个交易的最初外部账户

为什么不用 tx.origin 作为鉴权:
防止钓鱼攻击:
攻击者部署中间，诱导原始用户A，调用 转账功能。
```

## 三明治交易
```
三明治攻击是攻击者监听 mempool（未打包交易池）中即将执行的大额交易，
在 大概买入单 前后 夹入 买单 和 卖单 ，进行套利

mempool:当用户发送交易时，交易 首先到达节点的内存池（mempool），交易在这里等待矿工或验证者打包进区块；

防御:拆分大单交易，避免滑点过大
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
receive():接收 ETH ；
fallback():处理未匹配函数或 ETH 调用。
```


## payable
```
payable函数是一种特殊类型的函数，允许合约接受以太币作为支付
payable(address)
```


## Solidity字节码组成
```
Creation Code（部署字节码）
Runtime Code（运行字节码） :

如果部署一个空合约，部署字节码 和运行字节码 （STOP 指令） 都是存在的，只是很短。整个合约没有任何逻辑

```

## private 修饰的变量，可以背访问吗 
```
private 只是防止其他合约直接访问
但是，因为区块链是公开账本，storage 数据都是公开的，可以调用 web3.eth.getStorageAt  读取
web3.eth.getStorageAt(contractAddress, slotIndex)

解决:不要在合约中存储敏感信息，密码之类的。
```

## 为什么什么uint256 和 int256 存储 占用的字节都是 32 ，符号不占位置吗
```
使用 二进制补码（two’s complement） 表示符号
最高位（bit 255）是符号位
0 = 正数
1 = 负数
其余 255 位存储数值
```

## alldata 中的负数会消耗更多的 gas
```
先:填充到 32 字节（256 位）
再:最高位补 0
存储 消耗更多的gas
```








