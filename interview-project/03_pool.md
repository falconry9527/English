# clearpool 简介
```
clearpool 是一个去中心化借贷协议。
流动性提供者，将资产存入流动性池（Reserve）赚取利息；
借款人，通过质押资产进行借贷，并支付利息。
```

# 借贷协议主要方法
```
借贷协议主要方法 ：
poolCreation：创建流动性池子（Pool）
lending / supply：出借资产到协议，增加流动性，获得利息。
redeem：出借方赎回存款和利息 。
staking：质押资产到协议或流动性池子。
borrow：借款方从协议借入资产。
repayBorrow：借款方归还借款，包括本金和利息。
liquidation（清算）：当借款方抵押率低于阈值时，协议触发清算。
```

# poolCreation ： 创建池子
```
function initReserve(
    address asset,            // 要添加的ERC20资产地址
    address interestRateStrategyAddress, // 利率策略合约
) external;
```

# lending / supply ： 添加流动性
```
function deposit(
    address asset,        // 要存入的ERC20代币地址
    uint256 amount,       // 存入数量
    address onBehalfOf,   // 资产接收者地址（可以是自己或其他人）
) external;
```

# redeem：出借方赎回存款和利息 。
```
function withdraw(
address asset,         // 要赎回的资产地址（底层 ERC20）
uint256 amount,        // 赎回数量（aToken数量换算成底层资产）
address to             // 接收资产的地址
) external returns (uint256);
```

# staking ： 质押资产到协议或流动性池子
```
function stake(
    address onBehalfOf,  // 收益接收者地址
    uint256 amount       // 质押 AAVE 数量
) external;
```

# borrow ：借款
```
function borrow(
    address asset,           // 借款的ERC20资产地址
    uint256 amount,          // 借款数量
    uint256 interestRateMode,// 利率模式：1 = 稳定利率，2 = 可变利率
    address onBehalfOf       // 借款接收地址（可以是自己或其他人）
) external;
```

# repayBorrow：借款方归还借款，包括本金和利息。
```
function repay(
address asset,            // 借款资产地址
uint256 amount,           // 偿还数量（可以是部分或全部）
uint256 rateMode,         // 借款利率模式：1 = 稳定利率，2 = 可变利率
address onBehalfOf        // 还款目标地址（通常是借款人自己）
) external returns (uint256);
```


## 借贷协议利息计算
```
池子里每个存款代币都对应一个 累计利息
用户本次改变账户的利息 = 用户在借贷池里存入的资产/借贷池子中总的 × （新累计利息 - 上次更新的累计利息）
利息的增长 是持续进行的，而 利息的计算和更新 主要是在用户 偿还借款时（repay）或每个区块更新时进行的

累计利息和用户利息都是 不是实时更新的
累计利息： 定时批量更新（每一分钟，Aave是监控区块，每个区块更新一次 ）
用户利息 ：1.定时批量更新 2.用户查看的时候更新
updateBorrowInterest :
updateLendInterest : 

```





