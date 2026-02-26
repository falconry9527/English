
##  adaptive curve irm (自适应曲线利息模型) 
```
一.直线利率模型
end_rate_at_target = start_rate_at_target * speed * seconds
speed = ADJUSTMENT_SPEED *  err （池子的资产利用率偏差)

池子的资产利用率偏差 err:
池子的资产利用率 > 90% , 鼓励存款，加大利息 （ err ∈（0，1） ）
池子的资产利用率 < 90% , 鼓励借款，减小利息 （ err ∈（-1.0） ）

二. 曲线利率模型
linear_adaptation = speed * seconds

在 AdaptiveCurveIRM 里把 linear_adaptation 变成  （e^（linear_adaptation/WAD））* WAD, 变 直线利率 为 复利

泰勒近似值: e^r ≈ 1 + r+ x^2/2 ,当r很小 : e^r≈1+r
梯形积分取平均数
倍率缩放，提高即时反馈

``` 


##  份额失真问题
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
