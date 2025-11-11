## uniswap v3 的主要更新：
```
1. 池子数量和结构
V2：每个交易对只有一个池子。
V3：每个交易对可有多个池子（按手续费等级区分），每个池子管理不同价格区间的流动性。

2. 流动性提供方式
V2：LP 必须提供等值的两种代币，流动性在整个价格区间均有效。
V3：LP 可以选择指定价格区间（tickLower、tickUpper）提供流动性，支持单边流动性，资本效率更高。

3. 手续费（Fee）
V2：固定 0.3%，全部归 LP。
V3：支持多档手续费等级（0.05% / 0.3% / 1%），协议收取部分手续费。
```

## uniswap 的基本原理
```
Uniswap 是 去中心化交易所（DEX），
使用 AMM (Automated Market Maker) 自动做市商的方式  。

中心化交易所:      订单簿（Order Book）+ 撮合系统(最高买价 碰撞 最低卖价) 的自动交易机制。
Uniswap : AMM 通过 流动性池（Liquidity Pool）和 恒定乘积公式(x * y = k） 决定价格，并完成交易。

流动性池（Liquidity Pools）: 存入一个交易对(token0和token1) 资产的池子（如 ETH 和 USDT）。
流动性提供者（LP）， 向池子 中存入 资产，并 获取 手续费收益 的用户，  。
交易者 , 从流动性池子中，购买资产，并支付手续费

恒定乘积公式L: x * y = k :
X ：token0 的 数量
Y : token1 的 数量

定价公式: 
token0(对token1）的价格=Y/X

输出计算: 假如输入 Xa，根据公式
(x+xa) * (y-ya) = x*y  
计算出：
Ya=(Y- (X * Y) / (X+Xa) )

```

## 对于 原生资产ETH,没有合约地址，怎么处理的
```
token0 和 token1: 交易对的两种代币地址。
fee:交易池的手续费级别（例如:0.05%、0.30%、1%）。

1. 对于 原生资产 ETH,没有合约地址，怎么处理的
在 Uniswap V3 中，如果交易对涉及 ETH，ETH 会通过 WETH（Wrapped ETH）合约进行包装，作为 ERC-20 代币 使用。
```

## uniswap 的主要合约
```
PoolManager :  pool creation ， get 等
PositionManager : mint , burn , swap 
ISwapRouter : mint , burn , swapIn，swapOut

```

## uniswap 的核心逻辑
```
0.  pool creation ： 创建池子
1. mint: 添加流动性
2. burn: 移除流动性
3. swap: 交易
4. collect : 提取账户资产

tickLower 和 tickUpper:流动性提供者选择的价格区间（以 tick 的形式表示）。
sqrtPriceX96（可选）:池子当前的价格的平方根（对于初始创建来说不一定必须）。
```

## pool creation
```
function createPool(
    address token0,
    address token1,
    uint24 fee,
) external returns (address pool);

token0 和  token1 : 交易对
fee: 以 100万为基数： 500，3000 ，10000
fee:交易池的手续费级别（例如:0.05%、0.30%、1%）。

```

## mint : 添加流动性
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

## mint : 获取流动性： 流动性计算公式 
```
@param sqrtRatioX96 : 当前价格
@param sqrtRatioAX96 : 最高价
@param sqrtRatioBX96 : 最低价
@param amount0 : token0的数量
@param amount1 : token1的数量
@return liquidity : 流动性数值

sqrtRatioBX96 和 sqrtRatioAX96 是根据 tickLower，tickUpper 算出来的

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

## swap ：LP 获得的手续费（ lpUserFee）
```
一. 计算公式:
LP 获得的手续费 = 交易手续费 × （LP提供的流动性 ÷ 池子总流动性）

二. 过程  : 
--------手续费指数
核心: 手续费指数= 每单位流动性累计手续费= 交易手续费/池子总流动性

1.交易池子维护了一个 池子的手续费指数，每次更新 :
totalFee=totalFeeOld + Fee/ 池子总流动性

手续费指数, 避免了池子流动性变化的干扰

2. 流动性提供者 维护了一个 用户的手续费指数，每次更新 :
a. LP新增的手续费 = （池子最新 手续费指数totalFee- 用户 手续费指数的userFee） * LP提供的流动性
b. 更新 userFee=totalFee

更新策略：
1.定时更新
2.触发更新 : 查看账户信息 或者 移除流动性的时候

```

## uniswap 怎么防止滑点过大
```
1. 用户设置 minAmountOut，售出价格 如果小于 minAmountOut 则回滚
2. 代码设置滑点容忍度，滑点过大则回滚
```

## 数据同步过程中,如何处理区块链分叉问题？
```
处理方法
1. 等待 12 个区块确认（约 3 分钟）后再处理数据。
2. 同步 removed 字段，其中 removed = true 表示事件无效且已被废弃。

removed 字段是类似 blockNumber 的一个属性字段

由于区块已经被废弃，所以 同步 removed=true的数据方法
1. websocket 监控
2. 根据实时记录的blockhash ,进行循环查询

```