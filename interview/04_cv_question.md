## 数据同步过程中,如何处理区块链分叉问题？
```
硬分叉 (Hard Fork) ： 底层协议发生不兼容改动，旧节点无法理解新区块,长期分裂(Ethereum Classic)
软分叉 (Soft Fork) ： 协议规则收紧，旧节点仍能接受新区块,最终会收敛到一条链(比特币 SegWit)
短期分叉 (temporary Fork) ：由于网络延迟原因，进行的短期分叉,协议会自动选择最长链


处理方法
1. 等待 12 个区块确认（约 3 分钟）后再处理数据。
2. 同步 removed 字段，其中 removed = true 表示事件无效且已被废弃。

Q1 : how to sync on-chain event 
There are two ways to sync on-chain event 
1.WebSocket: Monitor on-chain events using WebSocket. 
2.HTTP: send HTTP requests in a loop to fetch data.

Q2 :How should we handle blockchain fork problem during data sync?
There are two ways to handle this problem. 
1. Wait for 12 block confirmations (about 3 minutes) before data sync.
2. Sync the removed field: removed field indicates whether the event was discarded during a fork.

discarded(dɪˈskɑːdɪd])
```

## 怎么防止交易滑点过大
```
Q3.How to prevent excessive slippage in trades
限价单(Limit Order)、价格预言机 (oracle )

excessive (ɪkˈsesɪv)
```

## Merkle tree
```
Merkle root :
merkle Proof : bytes32[]: 兄弟节点hash 路径（sibling hash path）
The data structure of BTC is a Merkle tree too.
Structure (ˈdeɪtə ˈstrʌktʃər)

Merkle Tree 是哈希树/二叉树：叶子节点存数据哈希，父节点存子节点哈希组合，根节点递归组合所有子节点哈希。
若任一子节点哈希变化，所有父节点哈希也会随之变化，从而便于高效验证。
对于白名单，存储时只需保存根节点哈希，无需保存所有用户地址，从而显著减少存储开销。
验证的时候，您只需要提供叶哈希和 Merkle 证明（兄弟哈希数组）。

Q4 : Please talk about Merkle Tree
A Merkle Tree is a hash tree/binary tree: 
leaf nodes store data hashes, 
parent nodes store combinations of child hashes, 
If any child hash changes, 
all parent hashes change too, enabling efficient verification.

For a whitelist, we only need to store the root hash instead of all user addresses, 
significantly reducing  costs of data storage . 

For verification, you only need to provide the leaf hash and 
the Merkle proof (which is an array of sibling node hashes).

significantly (sɪɡˈnɪfɪkəntli) 谁个‘泥肥cant理
enabling (ɪˈneɪblɪŋ)
sibling（ˈsɪblɪŋ）
verification (ˌvɛrəfəˈkeɪʃən)

```

## Collateralization Ratio（抵押率）
```
Q5 : Please talk about Collateralization
Collateralization Ratio =  Value  of  Collateralization / Value  of  Loan （200-300 %）
liquidation Threshold（ 清算阈值 150% )

Collateral （kəˈlætərəl）
Collateralization （kəlætərəlaɪ'zeɪʃn）靠-来特热li'zeɪʃn
Ratio （ˈreɪʃiəʊ）
Loan（loʊn）

Threshold（ˈθreʃhoʊld）

```

## 怎么保证投人的资产安全
```
Q6 : How to ensure the safety of investors' assets
1. 超额抵押（Over-Collateralization）
2. 清算机制（Liquidation Mechanism）

investors （ɪnˈvɛstərz）
```

##  闪电贷攻击（Flash Loan Attack）
```
1. 聚合多个价格源（Chainlink、备用 Oracle、DEX TWAP）
2. 在执行清算前再次验证
闪电贷攻击（Flash Loan Attack）
Oracle 数据安全（Chainlink, Band Protocol）
Oracle 价格操纵 (Price manipulation)

Q7 :   Please talk about Flash Loan Attack
Aggregate multiple price sources (Chainlink, backup oracles, DEX TWAP).
Validate the price again before executing liquidation.

price（praɪs）
prices（ˈpraɪsɪz）
Aggregate（ˈæɡrɪɡeɪt）

```


