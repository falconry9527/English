
##  adaptive curve irm (自适应曲线利息模型) 
```
一. 直线利率模型
linear_adaptation = speed * seconds
speed = ADJUSTMENT_SPEED *  err （池子的资产利用率偏差)

池子的资产利用率偏差 err(偏离因子):
池子的资产利用率 > 90% , 鼓励存款，加大利息 （ err ∈（0，1） ）
池子的资产利用率 < 90% , 鼓励借款，减小利息 （ err ∈（-1.0） ）

二. 曲线利率模型
1. 直线转曲线 
在 AdaptiveCurveIRM 里把 linear_adaptation 变成  （e^（linear_adaptation/WAD））* WAD, 变 直线利率 为 复利
end_rate = start_rate * （e^（linear_adaptation / WAD））* WAD

泰勒近似值: e^r ≈ 1 + r+ x^2/2 ,当r很小 : e^r≈1+r
梯形积分取平均数

2. cure  : 倍率缩放，提高即时反馈

``` 

##  清算奖励
``` 
每清算1个抵押物，可以获得的抵押物的数量
liquidation_factor = 1 /（1 - LIQUIDATION_CURSOR * (1 - lltv) ） 
LIQUIDATION_CURSOR = 0.3 

清算奖励 
liquidation_reward = liquidation_factor -1 

LLTV  清算奖励率
50%		15% 
80%     6.4%
90%     3.1%
95%     1.5%
``` 

##  虚拟份额
``` 
user_assets / user_shares = total_assets + VIRTUAL_ASSETS / total_shares + VIRTUAL_SHARES 

虚拟份额解决的问题:
1. shares / assets = 10^6  初始化值
2. 精度问题 
3. 攻击问题 （前提是允许捐赠，mropho使用的是记账，不是balanceof 所以不会有问题）

问题2:
如果 shares/assets  = 1 （初始化值）,那么当资产很少的时候，四舍五入的时候，会造成 shares/assets 失真

问题3:
1. 攻击者第一个存入 1 wei，获得 1 share（此时 totalShares=1, totalAssets=1）
2. 击者直接向合约捐赠 1e18 资产（不通过 supply，直接 transfer）
此时 totalShares=1, totalAssets=1e18+1
3. 受害者存入 1e18 资产，计算份额：1e18 × 1 / (1e18+1) = 0 shares（向下取整为 0）
4. 受害者得到 0 份额，但资产已被转入，全部归攻击者所有

``` 

##  精度问题阿里： 份额失真
``` 
用户资产 / 用户份额   =  total_assets/ total_shares
存款1：meta_vault_entry.deposit
+1028235685 ， +1028235685  -> （1028235825+140利息）= 1028235965 +1 / 1028235685 +1  =  1.0000001362 （ +1 是虚拟资产和份额）

取款2: meta_vault_entry.deposit
-500000000 ， -499999932   ->    (1028235825+2利息)=528235827 +1 / 528235753 +1 = 1.0000001363

取款2: meta_vault_entry.withdraw
-528235826 ，-528235753   ->    (1+1) / (0+1) = 2 

这是某个用户操作一个 vault ，一共有 1次存款，2次取款，取款2 操作之后，由于灰尘问题 ， 累计资产/ 累计份额 就失真了

但是，这种情况，只发生在资产被提取完的情况下，后续新存入资产，会按照新的比例初始化, 就是 资产 / 份额  的 比例感觉很怪，但是，不影响后续的 份额公式 

``` 