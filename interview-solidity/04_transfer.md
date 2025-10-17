# ETH 转账流程（面试讲解版）

## 1️⃣ 签名（Signing）

* 钱包/客户端准备交易字段：`nonce`、`to`、`value`、`gasLimit`、费用相关字段。
* 生成交易哈希（txHash），作为交易唯一 ID。
* 把交易发送给节点（存储 合约字节码 的节点）。


## 2️⃣ 验证（Validation）

* 节点收到交易，检查签名合法性，能恢复发送者地址。
* 检查 nonce 是否正确，防止交易顺序错乱或重放。
* 检查账户余额是否足够支付转账金额和 gas。
* 合法交易放入 mempool，准备被打包。

## 3️⃣ 打包（Inclusion in Block）

* 矿工或区块提议者从 mempool 选择交易（优先高费）。
* 填充交易到区块，并生成区块头，包含交易 Merkle root。
* 区块广播网络，等待共识确认（PoS 下最终性，PoW 下可能重组）。

## 4️⃣ 执行（Execution）

* EVM 按交易顺序执行：

    * 如果 `to` 是 EOA，直接余额更新。
    * 如果 `to` 是合约，执行合约代码，可能调用其他合约或触发事件。
* 若执行失败（revert），状态写入回滚，但 gas 已消耗。

## 5️⃣ 状态更新（State Update）

* 成功执行后，账户余额、nonce、合约 storage 被更新到全局状态树（Merkle‑Patricia Trie）。
* 交易收据生成，包括 `status`、`gasUsed`、logs 等。
* 区块头 `stateRoot` 更新，保证链上状态可验证。

## 6️⃣ 费用结算（Fee Settlement）

* 扣除实际消耗 gas 的费用：

    * EIP‑1559 下，`baseFee * gasUsed` 烧掉，`tip` 给打包者。
    * 老式 gasPrice 模型，全部费用给矿工。
* 未使用的 gas 部分退回给发送者（如果交易成功且 gas 足够）。

## 面试讲法总结

> “签名保证身份，节点验证合法性，矿工打包到区块，EVM 执行，状态更新到账户/合约，最后结算 gas。”

**关键词记忆法：**

* 签名 → 身份
* 验证 → 合法 / nonce / balance
* 打包 → mempool → block
* 执行 → EOA / 合约
* 状态 → 更新 / Trie
* 费用 → 消耗 / 退回 / 烧掉
