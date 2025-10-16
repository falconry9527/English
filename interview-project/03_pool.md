# clearpool 简介
```
clearpool 是一个去中心化借贷协议。
流动性提供者,将资产存入流动性池（Reserve）赚取利息；
借款人,通过质押资产进行借贷,并支付利息。
```

# 借贷协议主要方法
```
借贷协议主要方法 :
poolCreation :创建流动性池子（Pool）
lending / supply :出借资产到协议,增加流动性,获得利息。
redeem :出借方赎回存款和利息 。
staking :质押资产到协议或流动性池子。
borrow :借款方从协议借入资产。
repayBorrow :借款方归还借款,包括本金和利息。
liquidation（清算） :当借款方抵押率低于阈值时,协议触发清算。
```

# poolCreation : 创建池子
```
function initReserve(
    address asset,            // 要添加的ERC20资产地址
    address interestRateStrategyAddress, // 利率策略合约
) external;
```

# lending / supply : 添加流动性
```
function deposit(
    address asset,        // 要存入的ERC20代币地址
    uint256 amount,       // 存入数量
    address onBehalfOf,   // 资产接收者地址（可以是自己或其他人）
) external;
```

# redeem :出借方赎回存款和利息 。
```
function withdraw(
address asset,         // 要赎回的资产地址（底层 ERC20）
uint256 amount,        // 赎回数量（aToken数量换算成底层资产）
address to             // 接收资产的地址
) external returns (uint256);
```

# staking : 质押资产到协议或流动性池子
```
function stake(
    address onBehalfOf,  // 收益接收者地址
    uint256 amount       // 质押 AAVE 数量
) external;
```

# borrow :借款
```
function borrow(
    address asset,           // 借款的ERC20资产地址
    uint256 amount,          // 借款数量
    uint256 interestRateMode,// 利率模式 :1 = 稳定利率,2 = 可变利率
    address onBehalfOf       // 借款接收地址（可以是自己或其他人）
) external;
```

# repayBorrow :借款方归还借款,包括本金和利息。
```
function repay(
address asset,            // 借款资产地址
uint256 amount,           // 偿还数量（可以是部分或全部）
uint256 rateMode,         // 借款利率模式 :1 = 稳定利率,2 = 可变利率
address onBehalfOf        // 还款目标地址（通常是借款人自己）
) external returns (uint256);
```

## 借贷协议利息计算
```
一. 计算公式
出借方获得的利息 = 总利息 × （出借方存入的资产数量 ÷ 池子总资产数量）

二. 过程 :
核心：每单位资产累计利息 = 借贷产生的利息 ÷ 池子总资产数量
场景: 池子在一段时间内产生利息 Interest（池子总资产数量 × 利率 × 时间）

1. 定时更新 ： 池子每单位资产累计利息 :
池子最新每单位资产累计利息 totalInterest = 旧的每单位资产累计利息 totalInterestOld + Interest ÷ 池子总资产数量

每单位资产累计利息 累计手续费 , 避免了池子资产变化的干扰

2. 定时更新 ： 出借方新增的利息
a.出借方新增利息 = （池子最新每单位资产累计利息 totalInterest - 用户上次记录的每单位资产累计利息 userInterest） × 出借方存入的资产数量
b.更新用户的每单位资产累计利息记录：userInterest = totalInterest

更新策略：
1.定时更新
2.用户查看的时候更新
```


##  什么是闪电贷
```
闪电贷是一种 无抵押借贷机制，
允许用户在同一个区块内借出任意数量的资金，用于套利 等复杂操作。
只要在交易结束（即区块打包完成）前归还本金 +利息；否则，交易会原子性回滚，仿佛从未发生。

原子性 : 必须在同一区块中（10秒左右）完成 借出,套利, 归还本金和利息 操作。

闪电贷 成本 :
Gas 消耗：发起与执行操作时需支付的网络利息。
利息 ：通常约为 0.09%～0.3%，取决于协议（如 Aave、Uniswap、Balancer 等），在归还时一并支付 。

为什么闪电贷能被滥用？（根本原因）
无需前期资本：任何人都能瞬时获得巨大杠杆。
原子性：在单笔交易内完成一切，传统异步风控（人工/跨块检测）难以及时干预。

```

##  闪电贷攻击
```
闪电贷攻击是指攻击者利用 闪电贷 在同一笔原子交易内借入大量资金，进行操纵价格，清算与抵押套利 等攻击
案例: 
价格操纵（Oracle/AMM 操作）：在去中心化交易所（AMM）上用借来的资金大幅买入/卖出，短时间内改变池内价格， 利用 相反方向的加倍合约活力。
低价清算（Liquidation）：买入大量  借贷资产， 借贷资产 ，降低质押率 (质押资产/借贷资产 ) ，触发低价清算。

怎么防护:
1. 安全预言机机制 : Chainlink 使用 价格聚合机制 聚合多交易所价格，避免直接采用单一价格来源
2. 多源价格聚合 : 同时引入 多个预言机,AMM，Dex,Cex  交叉验证
3. 时间加权价格（TWAP）+ 延迟确认机制  :  增大时间窗口（30分钟），减少单次报价冲击影响。

```

## 怎么保证投人的资产安全
```
1. 超额抵押（Over-Collateralization）
2. 清算机制（Liquidation Mechanism）
抵押率= 质押的资产价值/借贷的资产价值 （200-300%）
清算阈值 : 150%
```

##  TiDB
```
TiDB 是一款兼容 MySQL 协议的分布式 MPP 数据库。
它支持 OLTP（联机事务处理） 和 OLAP（联机分析处理）。
```

