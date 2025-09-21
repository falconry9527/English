## 常见问题
```
storage、memory、calldata 的区别？
storage：链上永久存储，memory：函数内部临时存储，calldata：函数外部调用传入只读参数。

Storage, Memory, and Calldata Differences:
Storage: On-chain permanent storage.
Memory: Temporary storage used within functions.
Calldata: Read-only data passed as input from external function calls.

--------------------------
EVM的数据存储：
Memory (内存)     ：临时内存： 存储函数的参数和返回值。
Stack (栈)        ：临时值： 存储计算过程中的中间结果。
Storage (存储)    ：永久存储： 映射 (Mapping)，动态数组 等全局变量

EVM Data Storage:
Stack: Temporary values, storing intermediate results during computations.
Memory: Temporary memory, storing function parameters and return values.
Storage: Permanent storage, storing global variables such as mappings and dynamic arrays.

--------------------------
public、external、internal、private 区别？
public：链内外可访问；external：链外调用；internal：合约及继承可访问；private：仅当前合约可访问。

Solidity Visibility:
public: accessible internally and externally
external: callable only from outside
internal: accessible in contract and derived contracts
private: accessible only within the contract

--------------------------
view、pure、普通函数区别？
view 只读状态；pure 不读写状态；普通函数可修改状态。

Differences between view, pure, and regular functions in Solidity:
Solidity Function Types:
view: reads state only
pure: neither reads nor writes state
regular: can modify state

--------------------------
fallback 和 receive 的区别？
receive()：接收 ETH ；
fallback()：处理未匹配函数或 ETH 调用。

receive(): receives ETH
fallback(): handles unmatched function calls or receives ETH
```


## 基础语法与概念
```
构造函数是什么？
在合约部署时执行一次，用于初始化状态变量。

Constructor: executed once when the contract is deployed, used to initialize state variables.

--------------------------
delegatecall vs call
delegatecall 在调用者上下文执行，被调用合约修改调用者状态；call 修改被调用合约状态。

delegatecall vs call:
delegatecall: executes in the caller’s context, modifications affect the caller’s state.
call: executes in the callee’s context, modifications affect the callee’s state.

--------------------------
selfdestruct
销毁合约，返还剩余 ETH 给指定地址。

--------------------------
mapping 特点
无法遍历、默认值 0、常用于地址 → 数据。

Mapping Characteristics:
Cannot be iterated
Commonly used for address → data 
--------------------------
动态 vs 固定数组
动态可增减长度，固定长度 gas 更低。

Dynamic vs Fixed-size Arrays:
Dynamic arrays: length can increase or decrease
Fixed-size arrays: length is constant, lower gas cost
--------------------------
重入攻击 ： 解决：使用ReentrancyGuard。
权限控制 ：使用 Ownable 或 AccessControl 模块限制操作。

Reentrancy Attack:  mitigation: use ReentrancyGuard.
Access Control: restrict operations using Ownable or AccessControl modules.

```

## Difference between UUPS and Transparent Proxy
```
Transparent Proxy: 升级逻辑在代理合约，更费gas，更安全
UUPS ：升级逻辑在实现合约，更节省gas

Difference between UUPS and Transparent Proxy
Transparent Proxy: The upgrade logic resides in the proxy contract, which consumes more gas but is safer.
UUPS: The upgrade logic resides in the implementation contract, consuming less gas.
```

## 安全问题
```

```

## gas 优化
```

```


