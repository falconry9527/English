## EVM的数据存储
```
EVM 是运行在节点上的 虚拟机，每个节点一个 EVM 。 合约执行的节点（EVM）就是 存储合约字节码的节点 。

Memory：函数调用期间的临时存储，存放函数参数、局部变量，中间计算结果，调用结束后清空。
Stack：EVM 执行计算的核心区域，存储值类型的 局部变量，中间计算结果，生命周期极短，最多 1024 个 slot。
Storage：合约的永久链上状态存储，存放全局变量、映射（mapping）、可变数组、结构体等，操作成本高，需要消耗 Gas。
Calldata：外部函数调用的 只读输入参数存储区，存放 address、uint、bytes、string 等，生命周期仅在函数执行期间，成本低且不可修改。

全局变量：默认在 storage（包括 mapping、array、struct、值类型）
局部变量：
值类型 → 默认 stack
引用类型 → 必须指定 memory 或 storage（数组/struct）

mapping →  只能是 合约级全局变量（contract-level state),只能存储在 stroage,不能在函数中定义
1. Stack 存储的是 mapping 起始 slot 索引,指向具体的 stroage 的 start_slot
2. 每个元素的索引 :  keccak256(abi.encode(key, start_slot))

数组：
1. 栈上存储的是起始slot索引 : start_slot
2. 每个元素的索引： keccak256(start_slot) + index

```


## gas优化
```
一.常规优化
int 存储槽优化: 把小的变量定义在一起，拼凑成 32字节，打包到同一slot，减少storage的存储
使用 memory,calldata  存储  临时变量 ，以降低存储的gas 消耗。
uniswap 的tick 模式中：使用 位移运算，替代 乘以2 和 除以2 ；常用的中间结果 使用 全局变量存储，避免多次计算

二.代码架构的优化:
1. 复杂的统计逻辑，使用event 抓去做到链下进行统计
2. uniswap 单位流动性累计手续费 机制： 
3. 份额计算公式

```

##  安全和防御攻击
```
1. 重入攻击防护 :
a. 使用 ReentrancyGuard 合约 枷（类似于加锁）防止重入攻击。
b. 先检查,再修改状态, 最后 交互/资产转账（CEI: Check-Effect-Interaction）

ReentrancyGuard 会消耗额外 gas，主要是因为对 _status 变量的 storage 读写。
规范代码习惯: 不能随意调整变量的大小和顺序

2.闪电贷攻击和防护 : TWAP机制

3. 管理员权限
时间锁: 延迟一些危险操作，比如 修改费率
审核机制 : 
多签机制 : 防止密钥丢失或被盗: 
紧急暂停机制: 特殊情况下，暂停业务逻辑

```


##  UUPS 和 透明代理（Transparent Proxy）的相同和区别
```
设计目的 : 实现合约逻辑部分 可升级，同时保持状态不丢失。
模式相同 : 代理合约 + 实现合约
代理合约 : 接收用户调用，通过 delegatecall 转发逻辑； 数据存储（状态变量）。
实现合约 : 具体的业务逻辑 。

不同点:
Transparent: 代理合约管理升级，每次升级都要 权限检查逻辑+转发，更费gas，更安全
UUPS: 实现合约管理升级，代理只负责转发，轻量且更节省 Gas。

代码层级: 
openzepplin 更兼容的是: UUPS
TransparentUpgradeableProxy -> ProxyAdmin (权限管理)
UUPSUpgradeable -> Initializable(初始化)

合约升级的兼容性
1. 不能修改struct，但是可以新增
2. 不能修改函数的签名，但是可以修改函数的逻辑，也可以新增函数
```

## 升级合约的 存储冲突
```
代理合约 : 数据存储（状态变量），接收用户调用，通过 delegatecall 转发逻辑。
实现合约 : 具体的业务逻辑。

存储冲突: 
由于使用的是 delegatecall ，所以，实现合约的数据和修改 都是存储发生在在代理合约的

代理合约和逻辑合约的数据都存储在代理合约,两个和存储的 slot 初始化值（都是从0开始）和顺序相同， 那么就会数据相互覆盖，存储冲突 。

解决方法: EIP-1967
给代理合约的 变量 分配一个很大的，固定的slot初始化值，保证不会与逻辑合约的业务变量 slot 冲突:
uups: keccak256("eip1967.proxy.implementation") - 1;

storage slot 是 uint256 范围（≈ 1.1579×10⁷⁷），而 EIP-1967 _IMPLEMENTATION_SLOT ≈ 0.246×10⁷⁷
```


## delegatecall 和 call 的不同
```
不同一：
call 在被调用者上下文执行，修改被调用者状态。
例如：合约A调用合约B的函数，函数是在 B的上下文执行，修改b的变量

delegatecall 在调用者上下文执行，修改调用者状态。
例如： 可升级合约的代理合约，转发用户的函数给实现合约，函数是在 代理合约 的上下文执行，修改 代理合约 的变量 ；

不同二：
delegatecall: msg.sender 是实际调用者
call: msg.sender 是 直接调用者

用户A->合约B--delegatecall/call--> 合约C
合约C 的  msg.sender
delegatecall : 用户A
call : 合约B
```


## CREATE2
```
定义: 它们都是 EVM 中的合约创建指令（opcode）。

一. CREATE : 合约地址是不可预测的，地址由 部署者地址 + nonce 决定 
主要 依赖于部署顺序和次数（nonce） 
address = keccak256(rlp(sender_address, nonce))[12:]
sender_address：部署者地址
nonce：部署者账户的交易计数（每次交易 +1）

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
取决于 ：部署合约地址 + 盐值（salt） + 合约字节码

部署者地址（sender） → 部署合约地址（比如 Factory 合约地址）
盐值（salt） → 由部署者自定义（常用 keccak256(token0, token1)）
合约字节码（bytecode) → 待部署合约的字节码（creationCode）

deploy 使用 create2 真正部署(assembly 调用 CREATE2)；
getAddress 用相同参数计算地址；

三. 案例:
uniswap 的交易对池子是单独的合约，是唯一的，固定的。
使用 CREATE2 创建的，

四. 实现方式 
1. 使用 new 创建合约的时候，
不传 盐值（salt）,底层调用的就是 CREATE；
传入 盐值（salt）,底层调用的就是 CREATE2 
例如: new Pool{salt: salt}() 

2. assembly : 
assembly 关键字允许你 直接使用以太坊虚拟机（EVM）的低级指令（Yul/汇编），
绕过 Solidity 的高级抽象

```

## abi.encode 和 abi.encodePacked 之间有什么区别？
```
bytes memory data = abi.encode(a, b, c);
将输入参数按 ABI 编码（Solidity 标准 ABI）序列化，返回类型：bytes memory
可逆：每个参数都有固定大小（32字节填充）或 长度前缀
案例：函数调用、签名消息

不可逆：每个参数紧凑打包，没有固定大小（没有做32字节填充），长度前缀
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

## Merkle tree
```
数据结构:
Merkle Tree 是哈希树/二叉树:
叶子节点存数据哈希，
父节点存子节点哈希的组合哈希，
根节点存所有子节点哈希的组合哈希，
若任一子节点哈希变化，所有父节点哈希也会随之变化。

数据存储:
对于白名单，存储时只需保存根节点哈希，无需保存所有用户地址，从而显著减少存储开销。

数据验证:
验证的时候，您只需要提供叶验证哈希和 Merkle证明（叶验证哈希 兄弟节点哈希数组）。

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
payable 函数是一种特殊类型的函数，允许合约接受以太币作为支付
payable(address)

```


## Solidity字节码组成
```
Creation Code（部署字节码）
Runtime Code（运行字节码） ：

如果部署一个空合约，部署字节码 和运行字节码 （STOP 指令） 都是存在的，只是很短。整个合约没有任何逻辑

```

## private 修饰的变量，可以被访问吗 
```
private 只是防止其他合约直接访问
但是，因为区块链是公开账本，storage 数据都是公开的，可以调用 web3.eth.getStorageAt  读取
web3.eth.getStorageAt(contractAddress, slotIndex)

解决：不要在合约中存储敏感信息，密码之类的。
```

## 为什么什么uint256 和 int256 存储 占用的字节都是 32 ，符号不占位置吗
```
使用 二进制补码（two’s complement） 表示符号
最高位（bit 255）是符号位
0 = 正数
1 = 负数
其余 255 位存储数值
```

## calldata 中的负数会消耗更多的 gas
```
先：填充到 32 字节（256 位）
再：最高位补 0
存储 消耗更多的gas
```
