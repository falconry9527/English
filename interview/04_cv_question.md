## 数据抓取问题: 数据抓取的时候，怎么处理分叉
```
硬分叉 (Hard Fork) ： 底层协议发生不兼容改动，旧节点无法理解新区块,长期分裂(Ethereum Classic)
软分叉 (Soft Fork) ： 协议规则收紧，旧节点仍能接受新区块,最终会收敛到一条链(比特币 SegWit)
短期分叉 ：由于网络延迟原因，进行的短期分叉,协议会自动选择最长链

处理方法
1. 等待 12 个区块确认（约 3 分钟）
2. 根据 Removed = true 字断回滚数据
```

## 怎么防止交易 滑点过大
```
多池路由、限价单、价格预言机
```

## Merkle tree
```
Merkle root :
merkle Proof : bytes32[]: 兄弟节点hash 路径（sibling hash path）

验证一个地址是否在白名单中，只需要提供该地址及其对应的 Merkle proof（从叶子节点到根节点的兄弟节点哈希路径）
Verification of an address against the whitelist requires only the address and its  Merkle proof (sibling hash path).
```

