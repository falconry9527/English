#  Q1: block-chain:
```
Chain-based data structure: Each block is linked in chronological order.
Distributed ledger: Cryptography and consensus mechanism ensure data  immutability
chronological （ˌkrɒnəˈlɒdʒɪkl）（ˌkrɑːnəˈlɑːdʒɪkl）时间顺序的
ledger（ˈledʒər）
Cryptography （krɪpˈtɒɡrəfi）（krɪpˈtɑːɡrəfi）克瑞普-塔-格费
Consensus (kənˈsensəs)
immutability（ɪˌmjuːtəˈbɪləti）
ability(əˈbɪləti)
theory (ˈθiːəri)
```

#  Q2: Consensus Mechanism:
```
Before a data block is added to the chain,
it must be validated and confirmed by the majority of nodes (server nodes/miners).
Consensus Mechanism of BTC is Proof of Work ,
Consensus Mechanism of ETH is Proof of Stake ,
confirmed（kənˈfɜːrmd）
majority(məˈdʒɔːrəti)
major (ˈmeɪdʒər)
```

#  Q3: Proof of Work(Pow)
```
Miners compete for the right to add new blocks to the chain by performing extensive complex hash computations.
this process consumes a large  amount of energy.
compete (kəmˈpiːt)
extensive（ɪkˈstensɪv）
large（lɑːrdʒ）
consumes（kənˈsuːmz）
energy（ˈe/nərdʒi）
```

#  Q4 : Proof of Stake(PoS)
```
Validators compete for the right to add new blocks to the chain by staking cryptocurrency,
based on the amount and duration of their stake.
this process consumes far less energy than Proof of Work
```

#  Q5 : Bitcoin mining
```
Hash=SHA256(SHA256(Version+Prev Hash+Merkle Root+Timestamp+Bits+Nonce))<Target
Miners try different nonce values repeatedly（rɪˈpiːtɪdli）
and compute the double hash of the block header using SHA-256
until the hash is below the target,
earning the right to add the block.

The target value is adjusted based on both mining time and difficulty.
The purpose is to maintain an average block interval of about 10 minutes.

computation （ˌkɑːmpjuˈteɪʃn）
compute（kəmˈpjuːt）
average(ˈævərɪdʒ)
```

# Q6 : The data structure of the btc :
```
A Bitcoin block  contains two main parts:
---Block Header
Version: Block version number
Previous Block Hash: Hash of the previous block
Merkle Root: Root of the Merkle tree of all transactions in the block
Timestamp: Block creation time
Bits: Target difficulty
Nonce: Random number used for Proof of Work
---Transaction List
Contains all transactions packed in the block, each with its own inputs and outputs.
```

# Q7 : The data structure of the eth :
```
---Block Header
Parent Hash：父区块哈希
Uncle Hash：叔区块列表的哈希
Coinbase：出块矿工地址
State Root：状态树根
Transactions Root：交易树根
Receipts Root：收据树根
其他字段：Difficulty, Number, Gas Limit, Timestamp 等
---Transaction List
State Trie: Stores account states, including balances, nonces, and contract storage.
Transactions Trie: Stores all transactions in the block, ensuring transaction integrity.
Receipts Trie: Stores execution results and event logs for each transaction, used to verify outcomes.
```

# Q8 :  Difference between layer1,layer2 and layer3
```
L1 is the base（beɪs） blockchain, such as Bitcoin and Ethereum;
L2 is a scaling layer on top of L1.
It processes transactions off-chain in batches and submits the results back to the main chain
, which significantly increases TPS and reduces costs of gas , such as Polygon.
L3 is the application layer built on L2, such as Uniswap.
layer(ˈle/ɪə)
application(ˌæplɪˈkeɪʃn)
```

# Q9 : Cryptography
```
Symmetric Encryption: Encryption and decryption use the same key, but key transmission is not secure(e.g., HTTPS/TLS encryption).
Asymmetric Encryption: Encryption and decryption use a pair of different keys (public key + private key);
the public key is used for encryption or verification,
and the private key is used for decryption or signature(e.g., SSH).
Symmetric（sɪ'metrɪk）
Asymmetric(ˌeɪsɪˈmetrɪk)
Encryption(ɪnˈkrɪpʃ(ə)n)
Decryption(diˈkrɪpʃ(ə)n)
```

# Q10 :  blockchain Fork
```
Hard Fork:
An incompatible protocol upgrade.
Old nodes cannot validate new blocks.
It results in a permanent chain split.
Example: ETH → ETC.
Soft Fork:
A backward-compatible protocol upgrade.
Old nodes can still validate new blocks.
It eventually results in a single chain.
Example: Bitcoin SegWit.
```

# Q11 : P2P network
```
P2P network is a decentralized network
where nodes communicate and share data directly without a central server.
Each node acts as both a client and a server.
```