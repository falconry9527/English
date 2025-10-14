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

## 什么是智能合约
```
智能合约（Smart Contract）是区块链上的一种 自动执行的脚本 ，特点
条件触发自动执行：当满足预设条件时，合约会自动运行，无需人工干预。
不可篡改：一旦部署到区块链上，合约代码和执行结果无法被修改。
去中心化：运行在区块链网络上 。
```

## view、pure、普通函数区别？
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

