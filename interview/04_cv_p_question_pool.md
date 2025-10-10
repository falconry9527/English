## DAO
```
DAO（去中心化自治组织）是一种在区块链上通过智能合约自动执行规则、由社区成员共同治理的自治组织。
没有管理员
```

## evm 兼容链
```
1. Ethereum 用 ETH 支付 Gas，Polygon 用 MATIC，BSC 用 BNB。部署或交互时要使用对应代币。
2. Chainlink 等服务在不同链上地址不同，需要修改合约参数。
```

## 重入攻击防护
```
1. 使用 ReentrancyGuard 合约防止重入攻击。
2. 先检查和修改状态, 然后进行资产交易
3. 拉取支付（Pull Over Push）模式： 不直接向用户发送资金，而是让用户主动提取。
```

## ETH 三种转账方式
```
transfer 和 send 用于简单 ETH 转账，均有 2300 gas 限制；transfer 失败抛异常，send 失败返回 false。
call 是最灵活的方式，适用于复杂的合约交互（包括转账和函数调用），失败返回 false，但需要更多的错误处理。

transfer and send are used for simple ETH transfers, both with a 2300 gas limit; transfer throws an exception on failure, while send returns false on failure.
call is the most flexible method, suitable for complex contract interactions (including transfers and function calls), returns false on failure, but requires more error handling.
```

## 借贷协议利息计算
```
池子里每个存款代币都对应一个 累计利息
当池子流动性发生变化（用户借出、赎回、用户借入、归还借款）时，需要更新每单位代币的累计利息
用户本次改变账户的利息 = 用户在借贷池里存入的资产 × （新累计利息 - 上次更新的累计利息）

Each deposited token in the pool is associated with an accumulated interest 
When the pool's liquidity changes (users lending, redeem, borrow, or repayBorrow), the accumulated interest per unit of token is updated.
The interest earned on the user's account change = the user's deposited assets in the lending pool × (new accumulated interest - last updated accumulated interest)

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

Main methods of  lending protocol ：
PoolCreation: Create a liquidity pool.
Lending/Supply: Lend assets to the protocol to increase liquidity and earn interest.
Redeem: Lenders redeem their deposits and interest.
Staking: Staking assets to the protocol or liquidity pool.
Borrow: Borrow assets from the protocol.
RepayBorrow: Borrowers repay their loans, including principal and interest.
Liquidation: When a borrower's collateralization ratio falls below a threshold, the protocol triggers liquidation.
```

