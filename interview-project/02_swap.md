## uniswap 的基本原理
```
Uniswap 是 去中心化交易所（DEX），
使用 AMM (Automated Market Maker) 自动做市商的方式  。

中心化交易所:      订单簿（Order Book）+ 撮合系统(最高买价 碰撞 最低卖价) 的自动交易机制。
Uniswap : AMM 通过 流动性池（Liquidity Pool）和 恒定乘积公式(x * y = k） 决定价格，并完成交易。

流动性池（Liquidity Pools）: 存入一个交易对(token0和token1) 资产的池子（如 ETH 和 USDT）。
流动性提供者（LP）: 向池子 中存入 存入一个交易对 资产，并 获取交易手续费收入 的用户，  。

恒定乘积公式L: x * y = k :
X ：token0 的 数量
Y : token1 的 数量

定价公式: 
token0(对token1）的价格=Y/X
```

## 手续费范围
```
fee:交易池的手续费级别（例如:0.05%、0.30%、1%）。
```

## 对于 原生资产ETH,没有合约地址，怎么处理的
```
token0 和 token1: 交易对的两种代币地址。
fee:交易池的手续费级别（例如:0.05%、0.30%、1%）。

1. 对于 原生资产 ETH,没有合约地址，怎么处理的
在 Uniswap V3 中，如果交易对涉及 ETH，ETH 会通过 WETH（Wrapped ETH）合约进行包装，作为 ERC-20 代币 使用。
```

## uniswap 的核心逻辑
```
0.  pool creation ： 创建池子
1. mint: 添加流动性
2. burn: 移除流动性
3. swap: 交易

tickLower 和 tickUpper:流动性提供者选择的价格区间（以 tick 的形式表示）。
sqrtPriceX96（可选）:池子当前的价格的平方根（对于初始创建来说不一定必须）。

```

## pool creation
```
function createPool(
    address tokenA,
    address tokenB,
    uint24 fee
) external returns (address pool);
```


## mint 
```
address token0;
address token1;
uint256 tickLower
uint256 tickUpper;
address recipient;
uint256 deadline;

交易对地址: token0，token1
最高/最低流动性价格: tickLower , tickUpper
接收 流动性凭据的地址: recipient

```

## Uniswap V3 中的流动性计算公式 
```
@param sqrtRatioX96 : 当前价格
@param sqrtRatioAX96 : 最高价
@param sqrtRatioBX96 : 最低价
@param amount0 : token0的数量
@param amount1 : token1的数量
@return liquidity : 流动性数值

当前价格： P = token1 / token0
价格区间： [Pa, Pb] ，其中 Pa < P < Pb
内部采用平方根价格： √Pa, √P, √Pb

情况 1：仅提供 token0（当前价格低于区间）
L = (amount0 × √Pa × √Pb) / (√Pb - √Pa)

情况 2：仅提供 token1（当前价格高于区间）
L = amount1 / (√Pb - √Pa)

情况 3：同时提供 token0 与 token1（当前价格在区间内）
amount0 = L × (√Pb - √P) / (√P × √Pb)
amount1 = L × (√P - √Pa)

流动性计算公式 ，分三种情况
情况 1：仅提供 token0（当前价格低于区间）
情况 2：仅提供 token1（当前价格高于区间）
情况 3：同时提供 token0 与 token1（当前价格在区间内）

总的来说： 存入池子的代币数量amount0 /价格平方根的差 (√Pb - √Pa), 价格平方根的差 是一定的，主要取决于 存入池子的代币数量amount0

```

## mint：LP 获得的手续费（ lpUserFee）
```
一. 计算公式:
LP 获得的手续费 = 交易手续费 × （LP 提供的流动性 ÷ 池子总流动性）

二. 过程 (单位流动性 累计手续费 + 交替更新模式):

假如发生了一个交易：

交易产生 单位流动性 增长手续费 ：
单位流动性增长手续费 (growthFee) = 交易手续费 / 池子总流动性

交易前后池子 单位流动性 累计手续费 :
totalFee_before
totalFee_later=totalFee_before+growthFee

交易前后用户 单位流动性 累计手续费 :
userFee_before
userFee_later

1. userFeelast1 的价格更新为 totalFee_later
2. 计算  LP获得的手续费 并更新
 LP获得的手续费 = （userFee_later-userFee_before） * LP提供的流动性

```

## uniswap 怎么防止滑点过大
```
1. 用户设置 minAmountOut，售出价格 如果小于 minAmountOut 则回滚
2. 代码设置滑点容忍度，滑点过大则回滚
```

## uniswap 的闪电贷(可以先pass)
```
定义:Uniswap 的闪电贷（Flash Loan / Flash Swap）是一种允许用户在 无需预先资金的情况下 借入资产，
并要求 在同一笔交易内归还 的机制，常用于套利、清算或复杂 DeFi 操作。
无法在同一笔交易中归还本金 + 手续费，EVM 会立刻触发回滚，整个交易状态恢复到借贷前，资金池本金不会损失。

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
