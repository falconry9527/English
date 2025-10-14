## 区块链
```
区块链：
链式数据结构：每个区块按时间顺序依次链接形成链条。
分布式账本：通过密码学和共识机制，保持其不可篡改性 
```

## 共识机制
```
共识机制：数据块在加入链之前，必须得到大多数节点（服务器节点/矿机）的验证和确认。
工作量证明（Proof of Work, PoW）：矿工通过执行大量复杂的哈希计算来竞争记账权。这个过程消耗了大量的能量。
权益证明（PoS）：验证者通过质押的代币的数量和时间来获得记账权
```

## 挖矿（mining）
```
Hash=SHA256(SHA256(Version+Prev Hash+Merkle Root+Timestamp+Bits+Nonce))<Target

矿工们反复尝试不同的Nonce值，并使用SHA-256计算区块头的双重哈希，直到哈希值低于目标值，从而获得添加该区块的权利（记账权）。
目标值会根据挖矿时间和难度进行调整，以使平均出块时间保持在约10分钟左右。

```

## btc数据结构
```
btc 区块的数据结构：
A Bitcoin block  contains two main parts:
----Block Header
Version: Block version number
Previous Block Hash: Hash of the previous block
Merkle Root: Root of the Merkle tree of all transactions in the block
Timestamp: Block creation time
Bits: Target difficulty
Nonce: Random number used for Proof of Work

-----Transaction List
Contains all transactions packed in the block, each with its own inputs and outputs.
```

## eth数据结构
```
----区块头（Block Header）
Parent Hash：父区块哈希
Uncle Hash：叔区块列表的哈希
Coinbase：出块矿工地址
State Root：状态树根
Transactions Root：交易树根
Receipts Root：收据树根
其他字段：Difficulty, Number, Gas Limit, Timestamp 等

----交易列表（Transactions）

---- 三种树
状态树（State Trie）：存储账户状态，包括余额、nonce 和合约存储。
交易树（Transactions Trie）：存储区块内的所有交易，保证交易完整性。
收据树（Receipts Trie）：存储每笔交易的执行结果和事件日志，用于验证交易结果。
```

## layer1-3
```
L1 是基础主链，
L2 是 L1 之上的扩展层。 它在链下批量处理交易，并将结果提交回主链，从而显著提高 TPS 并降低成本，例如 Polygon。，
L3 是构建在 L2 上的应用层。
```

## 加密学
```
对称加密 ：加密和解密使用的是同一个密钥，但是，密钥传输不安全(e.g., HTTPS/TLS encryption) 。
非对称加密 ：加密和解密使用一对不同的密钥（公钥 + 私钥），公钥用于加密或验证，私钥用于解密或签名 (e.g., SSH)。
```

## 区块链分叉
``` 
硬分叉：
向后不兼容的协议升级。旧节点无法验证新区块。可能导致永久性的链分裂。示例：ETH → ETC。
a.调整区块大小限制（比如从 1MB 改为 2MB）
b.改变共识算法或挖矿难度计算方式

软分叉：
向后兼容的协议升级。旧节点仍然可以验证新区块。最终，该链会收敛为一条单链。示例：Bitcoin-SegWit。
a. 增加约束条件： 限制区块大小、限制交易格式

临时分叉：
由网络延迟或同时出块引起，短暂的链分裂，最终只保留最长链。

```

## P2P网络
``` 
P2P网络是一种去中心化的网络，节点之间无需中央服务器即可直接通信和共享数据。
每个节点既是客户端，也可以是服务器
``` 
