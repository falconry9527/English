## 架构
```
pool : 池子
position :  流动性头寸
tick : 
```

## pool_creator
```
pool_create_pool_v3 -> factory::create_pool_v3/create_pool_internal-> pool::new
create_pool_v3_with_creation_cap
区别 : is_permission_pair,有些交易对，只能拥有PoolCreationCap才能创建

唯一建 :   CoinTypeA，CoinTypeB , tick_spacing（fee_rate费率）

tick_lower_idx: u32,
tick_upper_idx: u32, 

get_sqrt_price_at_tick : 

```

## DEX 做市商
```
AMM  (Automated Market Maker) : 自动做市商  
CLMM (Concentrated Liquidity Market Maker) : 集中流动性做市商

中心化交易所: 订单簿（Order Book）+ 撮合系统(最高买价 碰撞 最低卖价) 的自动交易机制。
UniswapV2 : AMM 通过 流动性池（Liquidity Pool）和 恒定乘积公式(x * y = k） 决定价格，并完成交易。
UniswapV3 : CLMM 通过 流动性池（Liquidity Pool）和 恒定流动性公式(x * y = L^ 2 ; x = L / √P , y = L * √P） 决定价格，并完成交易。

流动性池（Liquidity Pools）: 存入一个交易对(token0和token1) 资产的池子（如 ETH 和 USDT）。
流动性提供者（LP）， 向池子 中存入 资产，并 获取 手续费收益 的用户,
交易者 , 从流动性池子中，购买资产，并支付手续费

无常损失:
无常损失是提供流动性时，价格偏离存入价导致 LP 资产价值低于直接持币的差额。
价格上涨或下跌都会造成 LP 价值低于持币，偏离越大损失越大。
LP 收益主要靠手续费和奖励抵消损失，高波动代币 HODL 更安全，高交易量低波动池提供流动性更有利。

```


## AMM 自动做市商  
```
恒定乘积公式L: x * y = k = L^2 
x : token0 的 数量
y : token1 的 数量

价格公式: 
token1/token0 的价格 p = y / x

交易swap公式（输出公式）: 
(x+Δx) * (y-Δy) = x * y  = k 

添加流动性公式(流动性计算公式):
ΔL = (Δx/x)  * L
ΔL = (Δy/y) * L

利息模型 : 
单位流动性累计利息 , 按照 流动性份额 分配

```


## CLMM : 集中流动性做市商
```
一 . 恒定流动性公式 :  
x * y = L^ 2 ; x = L / √P , y = L * √P

二. 价格公式 
初始价格 : √P = L / x  = y / L 
√P (池子会保留当前的价格 current_sqrt_price)

三. 交易公式 (输出公式) (L不变,√P变化,主要用于交易SWAP) 
Δx = L * (1/√P_new - 1/√P_old )
Δy = L * (√P_new  - √P_old )
循环更新一个tick, P_new ，按照上面公式， 算出 Δx_step，Δy_step 直到 sum(Δx_step) >= Δx , 算出 amount_in(Δx) 和 amount_out(Δy)

四. 流动性公式 (√P不变，ΔL 和 amount 变化，主要用于添加流动性)

CLMM 的三段状态 （输入amount0 和 amount1 可以获得多少流动）

情况 A：价格在区间内,同时提供 token0 和 token1 
1. token0 需求
amount0 = ΔL * (1/√P_cur - 1/√Pb)  ===>  ΔL= amount0 / (1/√P - 1/√Pb) 
2. token1 需求
amount1 = ΔL *  (√P_cur - √Pa)  ===>   ΔL= amount1 / (√P - √Pa) 

情况 B：价格低于区间，只需要 token0
amount0 =  ΔL * (1/√P - 1/√Pb) ===>  ΔL= amount0 / (1/√P - 1/√Pb)

情况 C：价格高于区间，只需要 token1
amount1 =  ΔL * (√P - √Pa)  ===>  ΔL= amount1 / (√P - √Pa) 

情况 B和C，池子退化成单边资产（但是池子其实还是有2种资产），只能添加单边资产

五. 流动性收益计算
1. 单位流动性累计利息
2. tick 反向存储(fee_below,fee_above) ,而不是存储落在当前tick的利息，读取的时候 fee_inside=fee_global−fee_below−fee_above

存储: 
1. 全局单位流动性累计利息
pool （update_fee_growth）
fee_growth_global_a
fee_growth_global_b

2. 全局单位流动性累计利息
Tick （pool.swap_in_pool-> 循环tick调用: tick.cross_by_swap ; 更新；increase_liquidity / decrease_liquidity -> update_by_liquidity 初始化）
fee_growth_outside_a
fee_growth_outside_b

3. 获取利息
领取 ：collect_fee-> get_fee_in_range
fee_inside=fee_global−fee_below−fee_above


```

## tick 价格模型
```
CLMM 的价格不是平滑的，而是每个价格差距 千分之一的 等差价格点

tick 与价格的关系是指数映射
p=1.0001^tick

1. 已知价格，求tick
tick=ln(1.0001)(p)

2. 已知tick求 sqrtPriceX64 
sqrtPriceX64 = √p * 2^64 = 1.0001^tick/2  *  2^64

代码gas优化 : 预计算常量 + 快速幂 + 移位运算

get_sqrt_price_at_tick : 根据 tick 求 sqrtPriceX64
get_tick_at_sqrt_price : 根据 sqrtPriceX64 求 tick

```
