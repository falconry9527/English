#  DEX 面试题 (中英双语)

## 一、基础概念 (Basic Concepts)
**Q1: 什么是 AMM（自动做市商）？ / What is AMM (Automated Market Maker)?**  
A: 一种基于算法公式的交易方式，不依赖订单簿。最常见的是恒定乘积公式 `x * y = k`（Uniswap V2）。  
A: A trading mechanism based on algorithmic formulas instead of order books. The most common is the constant product formula `x * y = k` (Uniswap V2).

**Q2: 什么是滑点 (Slippage)？ / What is Slippage?**  
A: 实际成交价格与预期价格的差异，主要来源于池子深度不足和交易量过大。  
A: The difference between the executed price and the expected price, mainly caused by insufficient liquidity depth or large trade size.

**Q3: 什么是无常损失 (Impermanent Loss)？ / What is Impermanent Loss?**  
A: 流动性提供者因价格波动导致的潜在损失，与直接持有资产相比，可能收益更低。  
A: A potential loss faced by liquidity providers due to price changes,
compared with holding the assets directly.

---


## 四、机制与性能 (Mechanism & Performance)
**Q10: 流动性提供者 (LP) 的收益从哪里来？ / Where do LP rewards come from?**  
A: 来自交易手续费分成，有时还有额外的流动性挖矿奖励。  
A: From transaction fees shared by traders, and sometimes extra liquidity mining incentives.

**Q11: Uniswap V2 vs V3 的区别？ / Difference between Uniswap V2 and V3?**
- V2: 全范围流动性，资本效率低。
- V3: 集中流动性，LP 可以选择价格区间，资本效率更高。
---  
- V2: Provides liquidity across the full range, lower capital efficiency.
- V3: Concentrated liquidity, LPs choose ranges, higher capital efficiency.

**Q12: Curve 为什么适合做稳定币交易？ / Why is Curve suitable for stablecoin swaps?**  
A: 采用 **稳定币优化曲线（恒定和 + 恒定积混合公式）**，在汇率接近 1:1 时滑点极低。  
A: Uses a **stablecoin-optimized curve (constant sum + constant product hybrid)**, offering very low slippage near 1:1 ratios.
---

## 五、行业与扩展 (Industry & Scaling)
**Q13: DEX 和 CEX 的区别？ / Difference between DEX and CEX?**
- DEX: 去中心化，资金自托管，交易透明，但速度和体验略差。
- CEX: 中心化，资金托管，撮合引擎高效，但存在信任风险。
---
- DEX: Decentralized, self-custody, transparent but slower and less user-friendly.
- CEX: Centralized, custodial, faster matching but trust-dependent.

**Q14: 什么是 DEX 聚合器？ / What is a DEX Aggregator?**  
A: 聚合多个 DEX 流动性（如 1inch、Matcha），为用户寻找最优交易路径。  
A: Aggregates liquidity from multiple DEXs (e.g., 1inch, Matcha) to find the best trade routes for users.

**Q15: Layer2 对 DEX 的影响？ / Impact of Layer2 on DEX?**  
A: 提高吞吐量、降低手续费，使得小额交易和高频交易更可行。  
A: Increases throughput and reduces fees, making small and high-frequency trades more feasible.  
