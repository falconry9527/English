
# Hyperliquid 的 EIP-712 签名

# 一句话
```
Hyperliquid 不发普通链上交易,而是把每个动作(下单/转账)
封装成 EIP-712 typed-data 让用户签名,签名 POST 给 /exchange 端点,
由 Hyperliquid L1 验签后执行。底层算法仍是 secp256k1 + ECDSA。
```

# 和普通 ETH 交易的不同(一句话)
```
普通 ETH 交易 = 签了就花 gas 立即上链执行;
EIP-712      = 
1.签一个结构化对象(字段化、钱包里可读、防钓鱼),
2.零gas不上链,
3.只是一张防重放的链下授权券,由合约/L1稍后凭券执行。

-> 所以 Hyperliquid 用 712 下单:既有私钥授权的安全和可读防钓鱼,
   又有交易所级的速度。
```

# 先分清两层(必背)
```
算法层(永远不变) : secp256k1 椭圆曲线 + ECDSA + Keccak-256 -> 产出 (r,s,v) 65字节
编码层(EIP-712)  : 决定"对什么数据去签",让钱包可读 + 防重放
EIP-712 不是新算法,只是"签之前怎么组织数据"的标准。
```

# EIP-712 的三段结构
```
domain  : 域分隔符 {name, version, chainId, verifyingContract} -> 防跨链/跨合约重放
types   : 字段类型声明(结构化)
message : 真正要签的业务数据
最终签名 = sign( keccak256( 0x1901 ‖ hashStruct(domain) ‖ hashStruct(message) ) )
```

# Hyperliquid 的两套 712(核心考点)
```
① L1 action 签名  —— 高频操作:下单 / 撤单 / 改单
   - 把 action 用 msgpack 编码 + nonce + vault地址 -> 哈希成 connectionId
   - connectionId 塞进固定结构 "Agent" 去签
   - domain = {name:"Exchange", version:"1", chainId:1337, verifyingContract:0x0...0}
   - chainId 固定 1337 = "phantom agent"设计
   - 钱包里基本不可读(只看到一个 connectionId hash),换来统一/低开销/高频

② User-signed action 签名 —— 敏感操作:转账 / 提现 / approve builder fee
   - 直接签可读字段(amount / destination / time ...)
   - domain.chainId = 真实链ID(如 Arbitrum 42161)
   - 钱包里能看到金额、目标地址 -> 可读、可审、防跨链重放
```

# 为什么要分两套
```
下单要快要频繁  -> phantom agent(L1),结构统一、签名内容不可读没关系(程序化高频)
动钱要安全可审  -> user-signed,字段可读,chainId 绑真实链,用户看懂自己转多少给谁
本质权衡: 高频性能 vs 资金安全可读性
```

# 在 seabond 项目里怎么落地
```
go-hyperliquid 的 Account 接口只要求一个方法:
    SignTypedData(ctx, td apitypes.TypedData) (*SignatureResult, error)

上层 Exchange 构造好 712 typed-data -> 调 SignTypedData -> 拿到 (r,s,v) -> POST /exchange
Exchange 完全不关心私钥在哪,只要能签出 EIP-712 签名。
-> 签名权与业务逻辑解耦,私钥永不进服务端代码。
```

两种签名托管后端(都实现同一个接口):
```
PrivyAccount  (account_privy.go) : 主账户,走 Privy 嵌入式钱包 RPC 签名
                                   注意 Privy 要 snake_case "primary_type" 而非标准 "primaryType"
VaultAccount  (account_vault.go) : 交易钱包,走 Vault gRPC 签名
                                   自己 HashStruct 算 domain separator(ds) + typed-data hash(tdh)
                                   再交给 Vault 用对应私钥签出 r/s/v
```

# 和 personal_sign / 普通交易的区别
```
eth_sendTransaction  : 真实上链交易(转ETH/调合约),会花 gas。Hyperliquid 下单不走这个。
personal_sign(191)   : 签一段加前缀的字符串,常用于登录,纯链下、不可结构化验证。
EIP-712(signTypedData_v4): 结构化、可读、可被合约/L1 解析验证。Hyperliquid 全程用它。
```

# 面试一句话
```
"Hyperliquid 全程用 EIP-712(底层 secp256k1/ECDSA),但分两套 domain:
 高频下单用 chainId=1337 的 phantom-agent 结构(把 msgpack action 哈希成 connectionId 再签,
 求统一和速度);转账提现这种敏感操作用真实 chainId 的 user-signed action,字段可读、防跨链重放。
 我们用 go-hyperliquid 的 Account 接口把签名抽象成 SignTypedData,
 私钥托管在 Privy 嵌入钱包或 Vault gRPC 里,服务端代码永远拿不到私钥,只拿签名结果。"
```

# 常见追问
```
Q: 为什么 L1 签名 chainId 是 1337?
A: 这是 Hyperliquid 的 phantom-agent 约定,固定值让所有 L1 action 用同一个 domain,
   签名内容是 action 的 hash(connectionId),不追求钱包可读,追求统一与吞吐。

Q: 下单签名会上链/花 gas 吗?
A: 签名动作本身不上链、不花 gas;签名 POST 给 Hyperliquid API,由其 L1 验签后撮合执行。

Q: nonce 在哪防重放?
A: action 里带 nonce(通常毫秒时间戳),L1 校验 nonce 防重放;
   user-signed 还额外靠真实 chainId + verifyingContract 防跨链跨合约重放。

Q: 私钥托管在哪?会不会被服务端拿到?
A: 不会。私钥在 Privy 嵌入钱包 / Vault 签名服务里,服务端只调 SignTypedData 拿回 (r,s,v)。
```

# 相关
[[envelope-encryption]] 信封加密 —— 同样是"密钥/签名权与业务解耦、私钥不出安全边界"的思路
