## 区块链
```
区块链：
链式数据结构：每个区块按时间顺序依次链接形成链条。
分布式账本：通过密码学和共识机制，保持其不可篡改性 

block-chain:
Chain-based data structure: Each block is linked in chronological order.
Distributed ledger: Cryptography and consensus mechanism ensure data  immutability

```

## 共识机制
```
共识机制：数据块在加入链之前，必须得到大多数节点（服务器节点/矿机）的验证和确认。

Consensus Mechanism:
Before a data block is added to the chain, 
it must be validated and confirmed by most nodes (server nodes/miners).
Consensus Mechanism of BTC is Proof of Work ,
Consensus Mechanism of ETH is Proof of Stake ,
```

## 工作量证明（Proof of Work, PoW）
```
中文：矿工通过执行大量复杂的哈希计算来竞争记账权。这个过程消耗了大量的能量。
Miners compete for the right to add new blocks to the chain by performing extensive complex hash computations.
this process consumes a large  amount of energy.

```

## 权益证明（Proof of Stake, PoS）
```
权益证明（PoS）：验证者通过质押的代币的数量和时间来获得记账权
English (Professional Version):
Proof of Stake (PoS): 
Validators compete for the right to add new blocks to the chain by staking cryptocurrency,
based on the amount and duration of their stake.
this process consumes far less energy than Proof of Work

```

## 挖矿（mining）
```
Hash=SHA256(SHA256(Version+Prev Hash+Merkle Root+Timestamp+Bits+Nonce))<Target

矿工们反复尝试不同的Nonce值，并使用SHA-256计算区块头的双重哈希，直到哈希值低于目标值，从而获得添加该区块的权利（记账权）。
目标值会根据挖矿时间和难度进行调整，以使平均出块时间保持在约10分钟左右。

English (Concise Version):
Miners repeatedly try different Nonce values 
and double-hash the block header with SHA-256 
until the hash is below the target, 
earning the right to add the block.
The target value is a difficulty parameter that the blockchain adjusts based on the actual mining time.
```

## btc数据结构
```
btc 区块的数据结构：
A Bitcoin block  contains two main parts:
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


## eth数据结构
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

State Trie: Stores account states, including balances, nonces, and contract storage.
Transactions Trie: Stores all transactions in the block, ensuring transaction integrity.
Receipts Trie: Stores execution results and event logs for each transaction, used to verify outcomes.
```

## layer1-3
```
L1 是基础主链，
L2 是 L1 之上的扩展层。 它在链下批量处理交易，并将结果提交回主链，从而显著提高 TPS 并降低成本，例如 Polygon。，
L3 是构建在 L2 上的应用层。
L1 is the base blockchain，such as Bitcoin and Ethereum
L2 is a scaling layer on top of L1.
It processes transactions off-chain in batches and submits the results back to the main chain
, which significantly increases TPS and reduces costs, such as Polygon.
L3 is the application layer built on L2, such as Uniswap.
```

## Cryptography
```
对称加密 ：加密和解密使用的是同一个密钥，密钥传输不安全(e.g., HTTPS/TLS encryption) 。
非对称加密 ：加密和解密使用一对不同的密钥（公钥 + 私钥），公钥用于加密或验证，私钥用于解密或签名 (e.g., SSH)。

Symmetric Encryption: Encryption and decryption use the same key, and key transmission is not secure(e.g., HTTPS/TLS encryption). 
Asymmetric Encryption: Encryption and decryption use a pair of different keys (public key + private key); 
the public key is used for encryption or verification, 
and the private key is used for decryption or signing(e.g., SSH).


```

