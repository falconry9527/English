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

记账权： 把新的区块添加到链上的权利 
```

## 共识机制:女巫攻击
```
定义：攻击者通过创建大量伪造身份（节点/账户），在去中心化网络中获取不成比例的影响力，从而操控网络行为。
PoW	： 每个身份需要付出算力成本，伪造大量身份代价极高
PoS	： 每个验证者需要质押资产，经济成本约束身份数量

```

## 共识机制:拜占庭将军问题
```
几位拜占庭将军各自率军围攻一座城池，他们必须通过信使通信来达成统一行动（进攻或撤退）。问题是：
信使可能被截杀（消息丢失）
某些将军可能是叛徒，故意发送矛盾的消息（恶意节点）

BFT（拜占庭容错）定理：在存在 f 个恶意节点的系统中，至少需要 3f + 1 个总节点才能达成共识。即系统最多容忍 不超过 1/3 的节点作恶。

PoW	: 用算力成本作为投票权，最长链为共识结果，作恶需要 >50% 算力	
PoS : 用质押资产作为投票权，作恶会被罚没（Slashing

```

## 挖矿（mining）
```
Hash=SHA256(SHA256(Version+Prev Hash+Merkle Root+Timestamp+Bits+Nonce))<Target

矿工们反复尝试不同的Nonce值，使用SHA-256,计算区块头的双重哈希，直到哈希值低于目标值，从而获得记账权（把新的区块添加到链上的权利 ）。
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
MPT 树 （Merkle Patricia Trie）: 是Merkle 的升级，存储键值对（address → account_state）
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

## DAO
```
DAO（去中心化自治组织）是一种在区块链上通过智能合约自动执行规则、由社区成员共同治理的自治组织。
没有管理员
```


## 双花问题
```
nonce + 余额检查 + 串行状态执行 + PoS 共识
nonce  : 这个账户已经成功发出的交易数量
```



## 零知识证明
```
证明者（Prover）向验证者（Verifier）证明某个陈述为真，但不泄露任何额外信息。

经典比喻：阿里巴巴洞穴
一个环形洞穴有左右两条路，中间有一道密码门。证明者想证明自己知道密码，但不想告诉验证者密码是什么：
证明者随机选一条路进洞穴
验证者在洞口喊"从左边出来"或"从右边出来"
如果证明者知道密码，无论被要求从哪边出来，都能做到（穿过密码门）
重复多次，如果每次都成功，验证者就相信证明者确实知道密码
全程验证者没有获得密码本身的任何信息。

零知识证明的核心价值是在不泄露数据的前提下证明数据的有效性，在区块链中主要用于扩容（ZK-Rollup）

```