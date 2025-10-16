## uniswap 的基本原理
```
1.Uniswap 是 去中心化交易所（DEX）

中心化交易所： 订单簿（Order Book）+ 订单撮合 的自动交易机制。
Uniswap : 使用 AMM (Automated Market Maker) 自动做市商的方式  。
通过 流动性池（Liquidity Pool）和 恒定乘积公式(x * y = k） 决定价格，并完成交易。

流动性池（Liquidity Pools）: 存入一对加密资产的池子（如 ETH 和 USDT）。
流动性提供者（LP）: 向池子 中存入 存入一对加密资产 获取 手续费收益 的用户 。

恒定乘积公式L: x * y = k :
X：tokenA 的 数量
Y: tokenA 的 数量

定价公式: 
TokenA 对 TokenB 的价格=Y/X

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

## 对于 原生资产ETH,没有合约地址，怎么处理的
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

## mint -> getLiquidityForAmounts
```
@param sqrtRatioX96 : 当前价格
@param sqrtRatioAX96 : 最高价
@param sqrtRatioBX96 : 最低价
@param amount0 : token0的数量
@param amount1 : token1的数量
@return liquidity : 流动性数值

liquidity=min(amount0/（sqrt(R_B)-sqrt(R_A)）,amount1/（sqrt(R_B)-sqrt(R_A)）)
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

## uniswap 怎么防止滑点过大
```
1. 用户设置 minAmountOut，售出价格 如果小于 minAmountOut 则回滚
2. 代码设置滑点容忍度，滑点过大则回滚
```

