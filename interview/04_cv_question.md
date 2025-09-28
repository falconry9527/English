## 数据抓取问题: 数据抓取的时候，怎么处理分叉
```
硬分叉 (Hard Fork) ： 底层协议发生不兼容改动，旧节点无法理解新区块,长期分裂(Ethereum Classic)
软分叉 (Soft Fork) ： 协议规则收紧，旧节点仍能接受新区块,最终会收敛到一条链(比特币 SegWit)
短期分叉 (temporary Fork) ：由于网络延迟原因，进行的短期分叉,协议会自动选择最长链

处理方法
1. 等待 12 个区块确认（约 3 分钟）
2. 根据 Removed = true 字断回滚数据

```

## 怎么防止交易 滑点过大
```
多池路由、限价单(Limit Order)、价格预言机
```

## Merkle tree
```
Merkle root :
merkle Proof : bytes32[]: 兄弟节点hash 路径（sibling hash path）

验证一个地址是否在白名单中，只需要提供该地址及其对应的 Merkle proof（从叶子节点到根节点的兄弟节点哈希路径）
Verification of an address against the whitelist requires only the address and its  Merkle proof (sibling hash path).
```

## Collateralization Ratio（抵押率）
```
Collateralization Ratio =  Value  of  Collateral / Value  of  Loan 
（200-300 %）

Liquidation Threshold（ 清算率 150% )

```

## 怎么保证投人的资产安全
```
How to ensure the safety of investors' assets
1. 超额抵押（Over-Collateralization）
2. 清算机制（Liquidation Mechanism）

```

##  闪电贷攻击（Flash Loan Attack）
```
闪电贷攻击（Flash Loan Attack）
Oracle 数据安全（Chainlink, Band Protocol）
Oracle 价格操纵 (Price manipulation)

1. 聚合多个价格源（Chainlink、备用 Oracle、DEX TWAP）
2. 在执行清算前再次验证

```
