## ETH 三种转账方式
```
transfer 和 send 用于简单 ETH 转账，均有 2300 gas 限制；transfer 失败抛异常，send 失败返回 false。
call  没有gas 限制，适用于复杂的合约交互（包括转账和函数调用），失败返回 false，但需要更多的错误处理。

transfer 和 send 都不推荐，原因：
固定 2300 gas 限制容易导致失败
受 EVM 升级影响，兼容性差
```

## 如何在 Solidity 中编写高效的 gas 循环？
```
1. 先读取storage ，缓存在memory，执行完毕之后，再写入 storage
```

## 代理合约中的存储冲突是什么
```
代理合约中的存储冲突是因为 delegatecall 在 Proxy 的存储上下文中执行 Logic 代码，
如果两者（代理合约和升级合约）声明的状态变量使用了相同 storage slot，
就会覆盖或破坏数据，需要使用标准槽位或非结构化存储模式避免冲突。
用于 Proxy 合约关键（assembly）变量的固定 storage slot 
```

## abi.encode 和 abi.encodePacked 之间有什么区别？
```
bytes memory data = abi.encode(a, b, c);
将输入参数按 ABI 编码（Solidity 标准 ABI）序列化
每个参数都有固定大小或长度前缀
返回类型：bytes memory
案例：函数调用、签名消息

将输入参数紧凑打包 不对齐、不加长度前缀
不可逆：解码时可能会出现二义性（例如两个动态类型参数拼在一起）
案例 ： 生成唯一标识、hash 计算
```

## uint8、uint32、uint64、uint128、uint256 都是有效的 uint 大小。还有其他的吗？
```
Solidity 支持 uint8, uint16, …, uint256，步长 8 位
uint256 = 256 bit = 2²⁵⁶ 个不同的值
```

## 在权益证明之前后，block.timestamp 发生了什么变化
```
PoW 阶段 block.timestamp 受矿工影响，波动较大；
PoS 阶段受协议严格限制，时间戳更稳定、可预测，区块间隔固定，操纵余地明显减少。
```

## OpenZeppelin ERC721 实现中的 safeMint 与 mint 有何不同？
```
mint 创建 NFT，safeMint 在创建的同时检查接收方是否支持 ERC721，确保 NFT 不会被锁死在合约中。
```

## Solidity 提供哪些关键字来测量时间？
```
block.timestamp
block.number
```

## 哪些操作会部分退还 gas？
```
CALL/DELEGATECALL 的失败 ： EVM 会退回剩余 gas
SELFDESTRUCT : 销毁合约
SSTORE（storage 置零）
```

## EVM 不支持浮点数指令
```
1. 精度问题
2. 和整数类型相比，会消耗更多gas
```

## 什么是 TWAP？
```
TWAP = 时间加权平均价格
核心优势：平滑价格，降低瞬时操纵风险
DeFi 中典型实现：
利用累计价格差 / 时间间隔计算
常用于抵押品估值、套利检测、交易策略

```



