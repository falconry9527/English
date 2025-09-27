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
derived (dɪˈraɪvd)

--------------------------
view、pure、普通函数区别？
view 只读状态；pure 不读写状态；普通函数可修改状态。

Differences between view, pure, and regular functions in Solidity:
view: can only read the state
pure: can neither read nor modify the state
regular: can both read and modify the state

pure (pjʊr)

--------------------------
fallback 和 receive 的区别？
receive()：接收 ETH ；
fallback()：处理未匹配函数或 ETH 调用。

receive(): receives ETH
fallback(): receives ETH or handles unmatched function calls 
```


## 基础语法与概念
```
构造函数是什么？
在合约部署时执行一次，用于初始化状态变量。

Constructor: executed once when the contract is deployed, used to initialize state variables.

initialize （ɪˈnɪʃəlaɪz）

--------------------------
delegatecall vs call
delegatecall 在调用者上下文执行，修改调用者状态；
call 在被调用者上下文执行，修改调用者状态；

delegatecall vs call:
call: executes in the callee’s context and modifies the callee’s state.
delegatecall: executes in the caller’s context and modifies  caller’s state.

callee(ˈkɔːli)
caller(ˈkɔːlə(r))
--------------------------
selfdestruct
销毁合约，返还剩余 ETH 给指定地址。

--------------------------
mapping 特点
无法遍历、默认值 0、常用于地址: 数据。

Mapping : Cannot be iterated, Commonly used for address: data 

iterated （ˈɪtəreɪtɪd）
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
UUPS: The upgrade logic is located in the implementation contract, which costs less gas.

located(ˈloʊkeɪtɪd)
```

## 安全问题
```
1. 权限控制: Ownable / AccessControl 
1. 多签:  TimelockController + AccessControl : 管理员提取费用需要至少 5 名管理员中的 3 名同意。
2. 时间锁: TimelockController 
3. 紧急停止: Pausable （ whenPaused/whenNotPaused）
4. 重入攻击-> ReentrancyGuard(nonReentrant)

security problem:
1. Access Control: Ownable / AccessControl
1. Multi-signature:TimelockController + AccessControl 
      Fees withdrawal by the admin require the consent of at least 3 out of 5 administrators.
2. Time-lock: TimelockController
3. Emergency Pause: Pausable 
4. Reentrancy Attack : use ReentrancyGuard.

```

## gas 优化
```
在存储中使用 uint256 代替较小的整数类型，以避免打包和解包的额外成本。
使用 calldata 作为外部函数参数，以降低 gas 成本。
使用 memory 作为临时变量，以避免存储写入。
逻辑优化：使用 Merkle 树作为白名单，以降低存储成本。

Use uint256 instead of smaller integer types in storage to avoid the extra costs of packing and unpacking.
Use calldata for external function parameters to reduce costs of gas .
Use memory for temporary variables to reduce  costs of storage/gas .
Logic Optimization: Use Merkle-tree–based whitelists to reduce costs of storage/gas  .

```


## OpenZeppelin
```
ReentrancyGuard  -> nonReentrant
Pausable -> whenPaused/whenNotPaused
Ownable/AccessControl -> onlyowner 

```


