## 数据同步过程中,如何处理区块链分叉问题？
```
处理方法
1. 等待 12 个区块确认（约 3 分钟）后再处理数据。
2. 同步 removed 字段，其中 removed = true 表示事件无效且已被废弃。
```

## evm 兼容链
```
1. Ethereum 用 ETH 支付 Gas，Polygon 用 MATIC，BSC 用 BNB。部署或交互时要使用对应代币。
2. Chainlink 等服务在不同链上地址不同，需要修改合约参数。
```

## Merkle tree
```
数据结构：
Merkle Tree 是哈希树/二叉树：
叶子节点存数据哈希，
父节点存子节点组合哈希，
根节点递归组合所有子节点哈希。
若任一子节点哈希变化，所有父节点哈希也会随之变化。

数据存储：
对于白名单，存储时只需保存根节点哈希，无需保存所有用户地址，从而显著减少存储开销。
数据验证：
验证的时候，您只需要提供叶哈希和 Merkle 证明（兄弟节点哈希数组）。
```

## DAO
```
DAO（去中心化自治组织）是一种在区块链上通过智能合约自动执行规则、由社区成员共同治理的自治组织。
没有管理员
```

## CREATE2
```
定义: 它们都是 EVM 中的合约创建指令（opcode）。

一. CREATE : 创建一个新合约，地址由 部署者地址 + nonce 决定 
主要 依赖于部署顺序和次数（nonce） ，不可预测
address = keccak256(rlp(sender_address, nonce))[12:]
sender_address：部署者地址
nonce：部署者账户的交易计数（每次交易 +1）

二. CREATE2 : 创建一个新合约，地址由部署者地址 + 合约字节码 + 盐值salt 决定
主要依赖于 盐值(salt: 手动传入的参数),所以，在部署前就能预测合约地址。
解决 主网与测试网部署地址不一致的问题；

计算方式:
address = keccak256(
    0xFF,
    sender,
    salt,
    keccak256(bytecode)
)[12:]  // 取后 20 字节
取决于 ： 部署者地址 + 盐值（salt） + 合约字节码

部署者地址（sender） → 部署者地址（比如 Factory 合约地址）
盐值（salt） → 由部署者自定义（常用 keccak256(token0, token1)）
合约字节码（bytecode) → 待部署合约的字节码（creationCode）

deploy 使用 create2 真正部署(assembly 调用 CREATE2)；
getAddress 用相同参数计算地址；

三. 案例:
iziswap 的 PoolFactory 使用 CREATE2 创建池子（Pair）：

四. new 
使用 new 创建合约的时候，
不传 盐值（salt）,底层调用的就是 CREATE；
传入 盐值（salt）,底层调用的就是 CREATE2 
例如: new Pool{salt: salt}() 
```

## transfer 和 safeTransfer
```
transfer 和 safeTransfer
transferFrom 和 safeTransferFrom

transfer ERC20 的转账方法，转账失败，返回false,不会回滚revert
function transfer(address to, uint256 amount) external returns (bool);

safeTransfer 是 OpenZeppelin对 transfer 的安全封装，
转账失败,返回false，并自动回滚revert, 保证转账安全。
SafeERC20.safeTransfer(IERC20 token, address to, uint256 value);

```

## view 函数会不会消耗gas 
```
不会: view函数虽然可以修改局部变量（会消耗gas），但是不修改全局变量，
gas 只会在 交易执行时（修改全局变量） 被实际消耗
```