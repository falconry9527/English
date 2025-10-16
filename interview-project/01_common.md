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
根节点递归两两组合所有子节点哈希。
若任一子节点哈希变化，所有父节点哈希也会随之变化。

数据存储：
对于白名单，存储时只需保存根节点哈希，无需保存所有用户地址，从而显著减少存储开销。
数据验证：
验证的时候，您只需要提供叶哈希和 Merkle 证明（兄弟节点哈希数组）。
```

## CREATE2
```
定义: 它们都是 EVM 中的合约创建指令（opcode）。

一. CREATE : 合约地址是不可预测的，地址由 部署者地址 + nonce 决定 
主要 依赖于部署顺序和次数（nonce） 
address = keccak256(rlp(sender_address, nonce))[12:]
sender_address：部署者地址
nonce：部署者账户的交易计数（每次交易 +1）

二. CREATE2 : 在部署前就能预测合约地址，地址由部署者地址 + 合约字节码 + 盐值salt 决定
主要依赖于 盐值(salt: 手动传入的参数) 。
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
uniswap 的交易对池子是单独的合约，是唯一的，固定的。
使用 CREATE2 创建的，

四. new 
使用 new 创建合约的时候，
不传 盐值（salt）,底层调用的就是 CREATE；
传入 盐值（salt）,底层调用的就是 CREATE2 
例如: new Pool{salt: salt}() 
```

## abi.encode 和 abi.encodePacked 之间有什么区别？
```
bytes memory data = abi.encode(a, b, c);
将输入参数按 ABI 编码（Solidity 标准 ABI）序列化，返回类型：bytes memory
不可逆：每个参数都有固定大小或长度前缀
案例：函数调用、签名消息

不可逆：每个参数紧凑打包，没有固定大小或长度前缀
解码时可能会出现二义性（例如两个动态类型参数拼在一起）
案例 ： 生成唯一标识
uniswap 的交易对池子是单独的合约，是唯一的，固定的 。
使用 encodePacked，根据交易对的地址，生成唯一标识-盐值

```

## Solidity 函数选择器
```
函数选择器是 Solidity 函数的唯一标识 。
它是 函数签名 functionName(type1,type2,...) 的 Keccak-256 哈希 的前 4 个字节。
bytes4 selector = bytes4(keccak256("transfer(address,uint256)"));

主要用于底层函数 call 和 delegatecall 的调用

4字节存储（232≈4.3×109 43 亿个值）

```

## transfer 和 safeTransfer
```
transfer 和 safeTransfer
transferFrom 和 safeTransferFrom

transfer ERC20 的转账方法，转账失败，返回false,不会回滚revert
function transfer(address to, uint256 amount) external returns (bool);

safeTransfer 是 OpenZeppelin对 transfer 的安全封装，
添加了 转账失败,返回false，自动回滚revert 的逻辑, 保证转账安全。
SafeERC20.safeTransfer(IERC20 token, address to, uint256 value);
```

## view 函数会不会消耗gas 
```
不会: view函数虽然可以修改局部变量（会消耗gas），但是不修改全局变量，
gas 只会在 交易执行时（修改全局变量） 被实际消耗
```

## 溢出 和下溢出
```
1. 溢出 ： 整数 超过 数据类型 的最大值  uint8 c = 156 + 1 ; // 溢出
1. 溢出 ： 整数 小于 数据类型 的最小值  uint8 c = a - b; // 下溢
unchecked 关键字用于 关闭整数运算的溢出/下溢检查 : 谨慎
```


## gas优化
```
一.常规优化
在存储中使用 uint256 代替较小的整数类型，多个小的整数类型，会被打包在一个slot，存在 打包和解包的资源消耗。
使用 memory, 存储 参数 和 临时变量，中间计算结果 ，以降低存储的gas 消耗。
使用 calldata 存储 外部函数 参数，以降低存储的gas 消耗。

二.代码架构的优化:
a. 链上只完成核心逻辑，复杂计算给到链下
比如: 贷款合约中，链下监控价格，触发链上清算，链上验证和执行

b. 贷款合约中，白名单机制使用 Merkle 而不是 mapping, 减少数据存储

c. iziswap中，定时计算 流动性提供者(LP)的手续费收入
每次swap ,流动性提供者(LP) 的手续费收入都会变化，如果每次计算 ，很消耗gas，怎么解决
1. 定时批量更新（Batching Updates）: 每10分钟更新一次
2. 当用户查看的时候更新

交易手续费 ： 平台抽取 1/6 , LP 获得 5/6 
0.05%（低风险池）
0.30%（中等波动池）
1.00%（高波动池）

```

##  安全和防御攻击
```
名单机制 :防止 Sybil 攻击

多签机制 ：防止密钥丢失或被盗: 
如果管理员想要提取交易费，需要至少 5 名管理员中的 3 名批准。
    使用多重签名机制，防止密钥丢失或被盗（OpenZeppelin 的 AccessControl + 自定义签名验证）。

重入攻击防护 :
1. 使用 ReentrancyGuard 合约防止重入攻击。
2. 先检查,再修改状态, 最后 交互/资产转账（Check-Effect-Interaction）
3. 拉取支付（Pull Over Push）模式: 不直接向用户转账，而是让用户主动提取。
   iziswap 中，流动性提供者（lp），赎回资产 ， 存入余额，让用户自己提取，而不是直接转账给账户

ReentrancyGuard 会消耗额外 gas，主要是因为对 _status 变量的 storage 读写。

规范代码习惯: 不能随意调整变量的大小和顺序
```

##  UUPS 和 透明代理（Transparent Proxy）的相同和区别
```
设计目的 : 实现合约逻辑部分 可升级，同时保持状态不丢失。
模式相同 : 代理合约 + 实现合约
代理合约 : 保存状态变量，接收用户调用，通过 delegatecall 转发逻辑。
实现合约 : 包含可执行代码。

不同点:
Transparent Proxy: 升级逻辑在代理合约，更费gas，更安全
UUPS :升级逻辑在实现合约，更节省gas

代码层级: 
openzepplin 更兼容的是: UUPS
TransparentUpgradeableProxy -> ProxyAdmin (权限管理)
UUPSUpgradeable -> Initializable(初始化)
```

## 什么是时间加权平均价格（TWAP）
```
TWAP = 时间加权平均价格，用于平滑价格波动。
核心作用：平滑价格、防止价格操纵、降低滑点。

Oracle：预言机的价格来源 
Uniswap: 大额订单执行(把大的订单拆分成很多小的订单，分时间段执行)。
```




