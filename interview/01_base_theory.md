## 区块链
```
区块链：
是一种链式数据结构：每个区块通过哈希指针连接到上一个区块，从而按时间顺序形成链。
是一种分布式账本：结合密码学哈希与共识机制，保证其不可篡改。

block-chain:
A chain-based data structure: each block is linked to the previous one through a hash pointer, forming a chronological chain.
A distributed ledger: immutability is ensured by cryptographic hashing and the consensus mechanism.
```

## 共识机制
```
共识机制：数据块在加入链之前，必须得到大多数节点（服务器节点/矿机）的验证和确认。

Consensus Mechanism:
Before a data block is added to the chain,
it must be validated and confirmed by the majority of nodes (server nodes/miners).

```

## 工作量证明（Proof of Work, PoW）
```
中文：矿工通过计算复杂哈希题竞争记账，第一个找到有效解的矿工可出块并获奖励，保证区块链安全与一致性。

English：Miners compete by solving complex hash puzzles; the first to find a valid solution can add a block and receive a reward, ensuring blockchain security and consistency.
```


## btc 区块的数据结构
```
btc 区块的数据结构：
A Bitcoin block consists of two main parts:
Block Header
Version: Block version number
Previous Block Hash: Hash of the previous block
Merkle Root: Root of the Merkle tree of all transactions in the block
Timestamp: Block creation time
Bits: Target difficulty
Nonce: Random number used for Proof of Work

Transaction List
Contains all transactions packed in the block, each with its own inputs and outputs.
```

## 挖矿
```
Hash=SHA256(SHA256(Version+Prev Hash+Merkle Root+Timestamp+Bits+Nonce))<Target

中文（简短版）：
比特币挖矿是不断尝试不同的 Nonce，使区块头的双重 SHA-256 哈希小于目标值，从而获得出块权和奖励。

English (Concise Version):
Bitcoin mining involves repeatedly trying different Nonce values until the block header’s double SHA-256 hash is below the target, 
earning the right to add a block and receive a reward.
```


## 权益证明（Proof of Stake, PoS）
```
中文（专业版）：
权益证明（PoS）：通过持有并质押代币来获得记账权，出块概率与质押数量成正比。

English (Professional Version):
Proof of Stake (PoS): Validators gain the right to propose blocks by holding and staking tokens, 
with the probability of being selected proportional to the amount staked.

```


## btc 区块的数据结构
```
区块头（Block Header）
Parent Hash：父区块哈希
Uncle Hash：叔区块列表的哈希
Coinbase：出块矿工地址
State Root：状态树根
Transactions Root：交易树根
Receipts Root：收据树根
其他字段：Difficulty, Number, Gas Limit, Timestamp 等

交易列表（Transactions）


状态树（State Trie）：存储账户状态，包括余额、nonce 和合约存储。
交易树（Transactions Trie）：存储区块内的所有交易，保证交易完整性。
收据树（Receipts Trie）：存储每笔交易的执行结果和事件日志，用于验证交易结果。

State tree: Stores account states, including balances, nonces, and contract storage.
Transactions tree: Stores all transactions in the block, ensuring transaction integrity.
Receipts tree: Stores execution results and event logs for each transaction, used to verify outcomes.
```
