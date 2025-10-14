## gas优化
```
一.常规优化
在存储中使用 uint256 代替较小的整数类型，以避免打包和解包的额外成本。
使用 memory, 存储 参数 和 临时变量 ，以降低存储的gas 消耗。
使用 calldata 存储 外部函数 参数，以降低存储的gas 消耗。

二.代码架构的优化:
a. 链上只完成核心逻辑，复杂计算给到链下（比如: 数据监控，数据统计 等）

b. 使用 Merkle 树作为白名单，以降低存储成本。

c. iziswap: 每次swap ,lp的userFee都会变化，如果每次计算 ，很消耗gas，怎么解决
1. 定时批量更新（Batching Updates）: 每10分钟更新一次
2. 当用户查看的时候更新

```

##  安全和防御攻击
```
女巫攻击 :使用基于 Merkle 树的白名单机制，防止 Sybil 攻击

防止密钥丢失或被盗: 
使用多重签名机制，防止密钥丢失或被盗（OpenZeppelin 的 AccessControl + 自定义签名验证）。
如果管理员想要提取交易费，需要至少 5 名管理员中的 3 名批准。
    
重入攻击防护 :
1. 使用 ReentrancyGuard 合约防止重入攻击。
2. 先检查,再修改状态, 最后 交互/资产转账（Check-Effect-Interaction）
3. 拉取支付（Pull Over Push）模式: 不直接向用户发送资金，而是让用户主动提取。
   swap burn : 存入余额，让用户自己提取，而不是直接转账给账户

规范代码习惯: 不能随意调整变量的大小和顺序
```

##  UUPS 和 透明代理（Transparent Proxy）的相同和区别
```
目的: 实现合约逻辑部分 可升级，同时保持状态不丢失。
模式相同 : 代理合约+ 实现合约
代理合约 : 保存状态变量，接收用户调用，通过 delegatecall 转发逻辑。
实现合约 : 包含可执行代码。

不同点:
Transparent Proxy: 升级逻辑在代理合约，更费gas，更安全
UUPS :升级逻辑在实现合约，更节省gas

代码层级: 
openzepplin 更兼容的是: UUPS
TransparentUpgradeableProxy -> ProxyAdmin (权限管理)
UUPSUpgradeable -> Initializable(初始化)

```

## 什么是 TWAP
```
TWAP = 时间加权平均价格
核心优势:平滑价格，降低瞬时操纵风险
DeFi 中典型实现:
利用累计价格差 / 时间间隔计算
常用于抵押品估值、套利检测、交易策略

```



