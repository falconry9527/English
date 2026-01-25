
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

在 AdaptiveCurveIRM 里把 linear_adaptation 变成  （e^（linear_adaptation/WAD））* WAD, 实现即时更新 利率的逻辑

泰勒近似值: e^r ≈ 1 + r+ x^2/2 ,当r很小 : e^r≈1+r
梯形积分取平均数
倍率缩放，提高即时反馈

``` 