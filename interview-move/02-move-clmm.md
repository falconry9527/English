## 架构
```
pool : 池子
position :  流动性头寸
tick : 
```

## pool_creator
```
pool_create_pool_v3
create_pool_v3_with_creation_cap
区别 : is_permission_pair,有些交易对，只能拥有PoolCreationCap才能创建

唯一建 :   CoinTypeA，CoinTypeB,tick_spacing（fee_rate费率）
tick_spacing :  fee_tiers: VecMap<u32, FeeTier>  

tick_lower_idx: u32,
tick_upper_idx: u32, 

get_sqrt_price_at_tick : 

```

## pool
```

```


## clmm_math : 流动性计算
```
get_liquidity_by_amount
get_amount_by_liquidity


当前价格： P = token1 / token0
价格区间： [Pa, Pb] ，其中 Pa < P < Pb
内部采用平方根价格： √Pa, √P, √Pb

get_liquidity_from_b :
情况 1：仅提供 token1（当前价格高于区间）
L = amount1 / (√Pb - √Pa)

get_liquidity_from_a :
情况 1：仅提供 token0（当前价格低于区间）
L = (amount0 × √Pa × √Pb) / (√Pb - √Pa)


```


## tick
```
```

