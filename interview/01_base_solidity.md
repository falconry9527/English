## 常见问题
```
--------------------------
EVM的数据存储：
Memory (内存)     ：临时存储： 存储函数的参数和返回值。
Stack (栈)        ：临时存储： 存储计算过程中的中间结果。
Storage (存储)    ：永久存储： 映射 (Mapping)，动态数组 等全局变量

EVM Data Storage:
Memory: Temporary storage, storing function parameters and return values .
Stack: Temporary storage, storing intermediate results during computations  within functions.
Storage: Permanent storage On-chain , storing global variables such as mappings and dynamic arrays.

Calldata: Read-only data , storing function parameters  from external calls.

intermediate(ˌɪntəˈmiːdiət)
results(rɪˈzʌlts) 
Permanent(ˈpɜːmənənt)
variables(ˈveəriəblz)
dynamic(daɪˈnæmɪk) 

--------------------------
public、external、internal、private 区别？
public：链内外可访问；
internal：合约及继承可访问；
private：仅当前合约可访问。
external：链外调用；

Solidity Visibility:
public: Accessible from anywhere: within the contract, derived contracts, and externally.
internal: Accessible within the contract and derived (child) contracts.
private:  accessible only within the contract
external: Accessible only from contract outside

Accessible （əkˈsesəbl）

--------------------------
view、pure、普通函数区别？
view 只读状态；pure 不读写状态；普通函数可修改状态。

Differences between view, pure, and regular functions in Solidity:
Solidity Function Types:
view: can only read the state
pure: can neither read nor modify the state
regular: can both read and modify the state

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
delegatecall 在调用者上下文执行，修改调用者状态；
call 在被调用者上下文执行，修改调用者状态；

delegatecall vs call:
call: executes in the callee’s context and modifies the callee’s state.
delegatecall: executes in the caller’s context and modifies  caller’s state.

--------------------------
selfdestruct
销毁合约，返还剩余 ETH 给指定地址。

--------------------------
mapping 特点
无法遍历、默认值 0、常用于地址: 数据。

Mapping :
Cannot be iterated, Commonly used for address: data 

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
Transparent Proxy: The upgrade logic is located in in the proxy contract, which costs more gas
UUPS: The upgrade logic is located in in the implementation contract, which costs less gas.

```

## 安全问题
```
1. 权限控制: Ownable / AccessControl 
1. 多签:  TimelockController + AccessControl : 管理员提取费用需要至少 5 名管理员中的 3 名同意。
2. 时间锁: TimelockController 
3. 紧急停止: Pausable （ whenPaused/whenNotPaused）
4. 重入攻击-> ReentrancyGuard(nonReentrant)

1. Access Control: Ownable / AccessControl
1. Multi-signature:TimelockController + AccessControl 
      Fee withdrawals by the admin require the consent of at least 3 out of 5 administrators.
2. Time-lock: TimelockController
3. Emergency Pause: Pausable 
4. Reentrancy Attack : use ReentrancyGuard.

```

## gas 优化
```
在存储数组中使用 uint256 而不是较小的类型，以避免昂贵的打包/解包操作。
使用 calldata 作为外部函数参数以节省 gas。
使用内存存储临时变量，避免存储写入操作。
逻辑优化：使用 Merkle 树作为白名单，以减少存储空间。

Use uint256 over smaller types in storage arrays to avoid costly packing/unpacking.
Use calldata for external function parameters to save gas.
Use memory for temporary  variables  avoid storage writes.
logic Optimizztion: use Merkle-tree for whitelists to reduce storage.
```


## OpenZeppelin
```
ReentrancyGuard  -> nonReentrant
Pausable -> whenPaused/whenNotPaused
Ownable/AccessControl -> onlyowner 

```


