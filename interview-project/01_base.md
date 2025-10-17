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
使用 encodePacked，uniswap中,生成交易对的唯一标识-盐值

```

## Solidity 函数选择器
```
函数选择器是 Solidity 函数的唯一标识 。
它是 函数签名 functionName(type1,type2,...) 的 Keccak-256 哈希 的前 4 个字节。
bytes4 selector = bytes4(keccak256("transfer(address,uint256)"));

主要用于底层函数 call 和 delegatecall 的调用

4字节存储（232≈4.3×109 43 亿个值）
```

