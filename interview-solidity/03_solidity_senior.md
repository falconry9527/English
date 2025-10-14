## 重入攻击防护
```
1. 使用 ReentrancyGuard 合约防止重入攻击。
2. 先检查和修改状态, 然后进行资产交易
3. 拉取支付（Pull Over Push）模式： 不直接向用户发送资金，而是让用户主动提取。
```

## gas优化
```
一.常规优化
在存储中使用 uint256 代替较小的整数类型，以避免打包和解包的额外成本。
使用 memory, 存储 参数 和 临时变量 ，以降低存储的gas 消耗。
使用 calldata 存储 外部函数参数，以降低存储的gas 消耗。

二.代码架构的优化：
a. 使用 Merkle 树作为白名单，以降低存储成本。
b. iziswap: 每次swap 都要算 每个用户的 userFee ，很消耗gas，怎么解决
1. 定时批量更新（Batching Updates）: 每10分钟更新一次
2. 当用户查看的时候更新

clearpool: 累计利息和用户利息都是 不是实时更新的
累计利息： 定时批量更新（每一分钟，Aave是监控区块，每个区块更新一次 ）
用户利息 ：1.定时批量更新 2.用户查看的时候更新
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
```


