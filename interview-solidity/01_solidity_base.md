## solidity 语言的数据类型
```
solidity 是一个静态语言，变量类型必须明确，专门用于以太坊智能合约开发

一. 值类型（Value Types）
1. 数字类型
整数（Integer）
有符号整数：int8 到 int256（步长为 8）
无符号整数：uint8 到 uint256（步长为 8）
默认：int 和 uint 等同于 int256 和 uint256

2.布尔值（Boolean）：值为 true 或 false

3.地址（Address）
address：存储 20 字节以太坊地址
address payable：可以接收 ETH 并执行 transfer 或 send，call

4. 字节和枚举
固定大小字节（Fixed-size byte array）
bytes1 到 bytes32
枚举（Enum）
enum Status { Pending, Shipped, Delivered }

二. 引用类型（Reference Types）
数组（Array）
固定长度数组：uint[5] numbers;
动态数组：uint[] numbers;
字节数组（Dynamic byte array）
字符串 : 本质是一个 byte 动态数组

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
动态数组：长度可变，元素是稀疏存储的，并不占用连续slot，一个元素占用一个slot（不管变量实际大小）
1. 栈上存储的是起始slot索引 : start_slot
2. 每个元素的索引： keccak256(start_slot) + index
slot连续，不代表物理存储连续
```

## mapping
```
mapping 是一种 键值对存储结构，底层使用的算法是 Hash：
1. 只能是 合约级全局变量（contract-level state),只能存储在 stroage,不能在函数中定义
2. 稀疏存储 ,元素是稀疏存储的
a. 栈上存储的是 mapping 起始 slot 索引： start_slot （指向具体的 stroage 的 slot）
b. 每个元素的索引 : keccak256(abi.encode(key,start_slot))

3. 不能迭代 如果要遍历，需要额外数组记录key
4. 删除 key
delete balances[addr] 会重置为默认值，但 slot 不回收
```

## interface
```
interface 是合约的抽象类型，只声明函数，不实现函数。
主要用于: 合约间交互,标准化接口,降低耦合
```

##  view、pure、普通函数区别？
```
view 只读状态；
pure 不读写状态；
普通函数可修改状态。
```

## view 函数会不会消耗gas
```
不会: view函数虽然可以修改局部变量（会消耗gas），但是不修改全局变量，
gas 只会在 交易执行时（修改全局变量） 被实际消耗
```

## 函数的可见性
```
public、external、internal、private 区别？
public：链内外可访问；
private：仅当前合约可访问。
internal：合约及继承可访问；
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
transfer 和 send 用于简单 ETH 转账，均有 2300 gas 限制；
transfer 失败抛异常，send 失败返回 false。
call  没有gas 限制，适用于复杂的合约交互（包括转账和函数调用），失败返回 false，但需要更多的错误处理。

transfer 和 send 都不推荐，原因：
固定 2300 gas 限制容易导致失败
受 EVM 升级影响，兼容性差
```

## transfer 和 safeTransfer
```
transfer 和 safeTransfer
transferFrom 和 safeTransferFrom

transfer ERC20 的转账方法，转账失败，返回false,不会回滚revert
function transfer(address to, uint256 amount) external returns (bool);

safeTransfer 是 OpenZeppelin对 transfer 的安全封装，
添加了 转账失败,返回false，自动回滚revert 的逻辑, 保证转账安全。
SafeERC20.safeTransfer(IERC20 token, address to, uint256 value);
```

## 溢出 和下溢出
```
1. 溢出 ： 整数 超过 数据类型 的最大值  uint8 c = 156 + 1 ; // 溢出
1. 溢出 ： 整数 小于 数据类型 的最小值  uint8 c = a - b; // 下溢
unchecked 关键字用于 关闭整数运算的溢出/下溢检查 : 谨慎
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

## gas
```
Gas 是 以太坊执行交易或智能合约操作的“燃料”。
交易费用=Gas× Price(一般以wei来计算，当前约 0.347 Gwei)
1 ETH = 10⁹ Gwei
1 ETH = 10¹⁸ Wei
```

## 智能合约大小大约可以有多大？
```
200kb
解决方法
1. 使用升级合约，减少每次升级的部署量
2. 模块化
3. 编辑器优化 optimizer
减少重复指令 → 减少字节码大小
合并常量和计算 → 减少部署 gas
```

## 什么是智能合约
```
智能合约（Smart Contract）是区块链上的一种 自动执行的脚本 ，
当满足预设条件时，合约会自动运行，无需人工干预。
极大丰富了区块链的应用生态
```


