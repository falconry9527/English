## go 语言的数据类型
```
一. 值类型（Value Types）
整数（Integer）
有符号整数：int8 到 int256（步长为 8）
无符号整数：uint8 到 uint256（步长为 8）
默认：int 和 uint 等同于 int256 和 uint256

布尔值（Boolean）：值为 true 或 false

地址（Address）
address：存储 20 字节以太坊地址
address payable：可以接收 ETH 并执行 transfer 或 send，call

固定大小字节（Fixed-size byte array）
bytes1 到 bytes32

枚举（Enum）
enum Status { Pending, Shipped, Delivered }

二. 引用类型（Reference Types）
数组（Array）
固定长度数组：uint[5] numbers;
动态数组：uint[] numbers;
字节数组（Dynamic byte array）

映射（Mapping）
键值对结构：mapping(keyType => valueType)
例子：mapping(address => uint) balances;

结构体（Struct）

三. 其他类型
函数类型（Function Types）
合约类型（Contract Types）
```

## 数组
```
定长数组：长度固定，元素是连续存储的，Gas 便宜
动态数组：长度可变，元素是稀疏存储的，适合链上数据可变场景

Stack 存储的是存储的指针 (slot索引) ,指向 storage的具体 slot:
keccak256(i) + index (i:初始化slot,index 元素的角标，一个元素占用一个slot)
```

## mapping
```
mapping 是一种 键值对存储结构，类似于 哈希表（HashMap）：
1. 稀疏存储
mapping 在 EVM 中是 稀疏存储，并不占用连续槽位，而是通过 哈希函数计算存储位置：
storage_slot=keccak256(abi.encodePacked(key,slot))
slot 是 mapping 声明时分配的槽位
key 是映射的键
keccak256 用于生成 唯一的存储位置
2. 不能迭代 如果要遍历，需要额外数组记录key
3. 删除 key
delete balances[addr] 会重置为默认值，但 slot 不回收
```

##  函数对状态变量 的操作性
```
view、pure、普通函数区别？
view 只读状态；
pure 不读写状态；
普通函数可修改状态。
```

## 函数的可见性
```
public、external、internal、private 区别？
public：链内外可访问；
internal：合约及继承可访问；
private：仅当前合约可访问。
external：链外调用；
```

## solidity的异常处理
```
require : 检查输入条件或调用前提，条件不满足则回退。
require(amount > 0, "Deposit must be positive");

revert : 回退交易并返回错误信息
revert("Insufficient balance");

assert : 只用于检测程序错误或不变量
uint c = a + b; assert(c >= a);  // 确保没有溢出

```

## ETH 三种转账方式
```
transfer 和 send 用于简单 ETH 转账，均有 2300 gas 限制；transfer 失败抛异常，send 失败返回 false。
call  没有gas 限制，适用于复杂的合约交互（包括转账和函数调用），失败返回 false，但需要更多的错误处理。

transfer 和 send 都不推荐，原因：
固定 2300 gas 限制容易导致失败
受 EVM 升级影响，兼容性差
```

## abi.encode 和 abi.encodePacked 之间有什么区别？
```
bytes memory data = abi.encode(a, b, c);
将输入参数按 ABI 编码（Solidity 标准 ABI）序列化
每个参数都有固定大小或长度前缀
返回类型：bytes memory
案例：函数调用、签名消息

将输入参数紧凑打包 不对齐、不加长度前缀
不可逆：解码时可能会出现二义性（例如两个动态类型参数拼在一起）
案例 ： 生成唯一标识、hash 计算
```


## OpenZeppelin ERC721 实现中的 safeMint 与 mint 有何不同？
```
mint 创建 NFT，safeMint 在创建的同时检查接收方是否支持 ERC721，确保 NFT 不会被锁死在合约中。
```

## Solidity 提供哪些关键字来测量时间？
```
block.timestamp
block.number
```

## 哪些操作会部分退还 gas？
```
CALL/DELEGATECALL 的失败 ： EVM 会退回剩余 gas
SELFDESTRUCT : 销毁合约
SSTORE（storage 置零）
```

## EVM 不支持浮点数指令
```
1. 精度问题
2. 和整数类型相比，会消耗更多gas
```

## 基础语法与概念
```
构造函数是什么？
在合约部署时执行一次，用于初始化状态变量。
```

## delegatecall 和 call 的不同
```
call 在被调用者上下文执行，修改被调用者状态。
例如：eth 转账 可以调用 cal 

delegatecall 在调用者上下文执行，修改调用者状态。
例如： 可升级的代理合约，用于转发逻辑；
```

## gas
```
Gas 是 以太坊执行交易或智能合约操作的“燃料”。
交易费用=Gas× Price(一般以wei来计算，当前约 0.347 Gwei)
1 ETH = 10⁹ Gwei
1 ETH = 10¹⁸ Wei
```


