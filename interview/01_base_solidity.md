## 常见问题
```
--------------------------
Storage (存储)    : 链上永久存储，因此操作它需要消耗 Gas，且非常昂贵： 映射 (Mapping)，动态数组 等全局变量


EVM的数据存储：
Memory：临时空间，用于函数调用时存储参数与返回值
Stack：EVM 的计算引擎，用于存储中间计算结果
Storage：链上永久存储，用于存储全局状态变量,比如mapping或者可变数组,所以，他消耗gas，而且，特别昂贵
Calldata：外部调用传入的只读参数区


EVM Data Storage:
Memory: Temporary Memory, stores function parameters and return values .
Stack:  The computation engine of the EVM, stores intermediate results during computations  within functions.
Storage: Permanent storage On-chain , stores global variables such as mappings and dynamic arrays.
  Therefore, it consumes gas and is very expensive.
Calldata: Read-only data , stores function parameters  from external calls.

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
constant（ˈkɑːnstənt）
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
安全问题：
1. 使用基于 Merkle 树的白名单机制，防止 Sybil 攻击。
2. 使用多重签名机制，防止密钥丢失或被盗（OpenZeppelin 的 AccessControl + 自定义签名验证）。
    如果管理员想要提取交易费，需要至少 5 名管理员中的 3 名批准。
3. 使用 OpenZeppelin 的安全合约（例如：Ownable、AccessControl、TimelockController、Pausable、ReentrancyGuard）。
a. 使用 Pausable 合约进行紧急暂停。
b. 使用 ReentrancyGuard 合约防止重入攻击。

Q8: security problem:
1. Use a whitelisting mechanism to prevent Sybil attacks.
2. Use a Multi-signature mechanism  to prevent the loss or theft of private key.
    If an admin wants to withdraw trading fees, they will need approval from at least 3 out of 5 administrators.
3.Use OpenZeppelin’s security contracts (Ownable, AccessControl,TimelockController,Pausable, ReentrancyGuard  ).
a.  Pausable contract is used to Emergency Pause
b.  ReentrancyGuard is used contract to prevent Reentrancy Attacks
   Update state first, then transfer funds

theft(θeft)
approval (əˈpruːvl)
funds(fʌndz)
Pause(pɔːz)
Pausable （ˈpɔː.zə.bəl）

contract (kənˈtrækt)
administrators(ədˈmɪnɪstreɪtəz)
Constructor (kənˈstrʌktər)
```

## 重入攻击防护
```
1. 使用 ReentrancyGuard 合约防止重入攻击。
2. 先检查和修改状态, 然后进行资产交易
3. 拉取支付（Pull Over Push）模式： 不直接向用户发送资金，而是让用户主动提取。
```

## gas 优化
```
在存储中使用 uint256 代替较小的整数类型，以避免打包和解包的额外成本。
使用 calldata 作为外部函数参数，以降低 gas 成本。
使用 memory 作为临时变量，以避免存储写入。
逻辑优化：使用 Merkle 树作为白名单，以降低存储成本。

iziswap: 每次swap 都要算 每个用户的 userFee ，很消耗gas，怎么解决
1. 定时批量更新（Batching Updates）: 每10分钟更新一次
2. 当用户查看的时候更新

clearpool: 累计利息和用户利息都是 不是实时更新的
累计利息： 定时批量更新（每一分钟，Aave是监控区块，每个区块更新一次 ）
用户利息 ：1.定时批量更新 2.用户查看的时候更新

gas optimization :
1.Use uint256 instead of smaller integer types in storage to avoid the extra costs of slot packing and unpacking.
2.Use calldata for external function parameters to reduce costs of gas .
3.Use memory for temporary variables to reduce  costs of storage/gas .
4.code logic Optimization: Use Merkle-tree–based whitelists instead of a mapping of addresses to reduce costs of storage/gas  .

```


## OpenZeppelin
```
ReentrancyGuard  -> nonReentrant
Pausable -> whenPaused/whenNotPaused
Ownable/AccessControl -> onlyowner 

```


