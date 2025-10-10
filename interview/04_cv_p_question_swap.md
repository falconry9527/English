## uniswap v1 ,v2 和v3 主要的区别
```
Uniswap V1:仅支持 ETH ↔ ERC20 交易，流动性提供者资本效率低。
Uniswap V2:支持 ERC20 ↔ ERC20 直接交易，并引入闪电贷和 on-chain 预言机。
Uniswap V3:引入集中流动性和多手续费等级，大幅提升资本效率和灵活性。
```

## uniswap 的闪电贷
```
定义:Uniswap 的闪电贷（Flash Loan / Flash Swap）是一种允许用户在 无需预先资金的情况下 借入资产，
并要求 在同一笔交易内归还 的机制，常用于套利、清算或复杂 DeFi 操作。

借出的资产存在托管账户 。
手续费: 一次性手续费（例如 Uniswap V2 通常为借款总额的 0.3%）
临界点: 无法在同一笔交易中归还本金 + 手续费，EVM 会立刻触发回滚，整个交易状态恢复到借贷前，资金池本金不会损失。

用户的风险: 
1. 风险较高L多重套利，用户的套利必须多余借款的利率，失败率较高
2. 如果失败，需要支付 gas ，复杂的套利需要消耗较多的gas

触发点:
1. Gas 费用必须提前支付
2. 触发是否回滚（无法在同一笔交易中归还本金 + 手续费），是在用户套利结束的时候回滚

避免闪电贷损失的方法
1. 仅在回报覆盖手续费和 本金 时执行操作 （0.3，那边借出的利率必须高于0.3）
2. 设置最大允许亏损阈值 → 如果潜在亏损过大，直接 revert 避免浪费 Gas
```

## createPool
```
token0 和 token1:交易对的两种代币地址。
fee:交易池的手续费级别（例如:0.05%、0.30%、1%）。
tickLower 和 tickUpper:流动性提供者选择的价格区间（以 tick 的形式表示）。
sqrtPriceX96（可选）:池子当前的价格的平方根（对于初始创建来说不一定必须）。

1. 对于 原生资产 ETH,没有合约地址，怎么处理的
在 Uniswap V3 中，如果交易对涉及 ETH，ETH 会通过 WETH（Wrapped ETH）合约进行包装，作为 ERC-20 代币 使用。
```

## mint 
```
address token0;
address token1;
uint256 amount0
uint256 amount1;
address recipient;
uint256 deadline;
```

## mint -> AMM
```
AMM（自动做市商，Automated Market Maker）
它通过 智能合约和流动性池 来实现资产的买卖 而不是订单薄 。
流动性池（Liquidity Pools）:AMM 基于流动性池来工作，流动性池是由两种或多种加密资产组成的池子（如 ETH 和 USDT）。这些资产由 流动性提供者（LPs） 存入池子中。
常数乘积公式（x * y = k）:AMM 使用数学公式来保持池子中的两种资产的平衡。例如，Uniswap 使用 x * y = k 的公式，其中:
k值在mint和burn(添加和减少流动性)的时候会成比例增大或缩小，在 swap的时候保持不变
```

## mint -> getLiquidityForAmounts
```
@param sqrtRatioX96 : 当前价格
@param sqrtRatioAX96 : 最高价
@param sqrtRatioBX96 : 最低价
@param amount0 : token0的数量
@param amount1 : token1的数量
@return liquidity : 流动性数值

liquidity=min( amount0/（sqrt(R_B)-sqrt(R_A)）,amount1/（sqrt(R_B)-sqrt(R_A)）)
amount0 和 amount1 分别是流动性提供者提供的两种资产的数量（例如，ETH 和 USDT）。
sqrt(R_A) 和 sqrt(R_B) 是价格区间的下限和上限的平方根，表示价格范围 [tickLower, tickUpper]。

当 ETH 的价格上升 时，burn（销毁）流动性 时提取的资产的 总价值（以 USDT 计）通常会增加
```

## mint -> userFee
```
userFee=(L_provided/L_total)×(totalFeeGrowth−feeGrowthInsideLastX128)
L_provided:流动性提供者在池子中的流动性份额。
L_total:池子中的总流动性。
feeGrowthInside:当前时刻该区间内的 手续费增长。
feeGrowthInsideLastX128:流动性提供者上次更新时的 手续费增长。

feeGrowthInside0LastX128 和 feeGrowthInside1LastX128 通常会在 swap（交易） 或 burn（销毁流动性） 时更新。
userFee也通常会在 swap（交易） 或 burn（销毁流动性） 时更新。

每次swap 都要算 每个用户的 userFee ，这样岂不是很消耗gas，怎么解决
1. 批量更新（Batching Updates）: 🇺每天更新一次
2. 当用户查看的时候更新

```