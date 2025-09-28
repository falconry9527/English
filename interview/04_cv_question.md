## 数据抓取问题: 数据抓取的时候，怎么处理分叉
```
硬分叉 (Hard Fork) ： 底层协议发生不兼容改动，旧节点无法理解新区块,长期分裂(Ethereum Classic)
软分叉 (Soft Fork) ： 协议规则收紧，旧节点仍能接受新区块,最终会收敛到一条链(比特币 SegWit)
短期分叉 (temporary Fork) ：由于网络延迟原因，进行的短期分叉,协议会自动选择最长链

处理方法
1. 等待 12 个区块确认（约 3 分钟）后再处理数据。
2. 根据 removed 字段回滚相关数据：若 removed = true，说明区块发生分叉，该事件已无效。

1. Wait for 12 block confirmations (≈3 minutes) before processing data.
2. Roll back data based on the removed field: if removed = true, the block was part of a fork and the event is invalid.

```

## 怎么防止交易 滑点过大
```
限价单(Limit Order)、价格预言机
```

## Merkle tree
```
Merkle root :
merkle Proof : bytes32[]: 兄弟节点hash 路径（sibling hash path）

Merkle Tree 是哈希树/二叉树：叶子节点存数据哈希，父节点存子节点哈希组合，根节点递归组合所有子节点哈希。
若任一子节点哈希变化，所有父节点哈希也会随之变化，从而便于高效验证。
A Merkle Tree is a hash tree/binary tree: 
leaf nodes store data hashes, 
parent nodes store combinations of child hashes, 
and the root node recursively combines all child hashes.
If any child hash changes, 
all parent hashes change accordingly, enabling efficient verification.

对于白名单，存储时只需保存根节点哈希，无需保存所有用户地址，从而显著减少存储开销。
For a whitelist, only the root hash needs to be stored, not all user addresses, 
significantly reducing costs data storage.

验证的时候，您只需要提供叶哈希和 Merkle 证明（兄弟哈希数组）。
For verification, you only need to provide the leaf hash and the Merkle proof (an array of sibling hashes).

```

## Collateralization Ratio（抵押率）
```
Collateralization Ratio =  Value  of  Collateral / Value  of  Loan 
（200-300 %）

Collateralization Ratio （抵押率）
LTV (Loan-to-Value Ratio) (和抵押率互为倒数)
Liquidation Threshold（ 清算阈值 150% )

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

Aggregate multiple price sources (Chainlink, backup oracles, DEX TWAP).
Re-validate prices before executing liquidation.

```


