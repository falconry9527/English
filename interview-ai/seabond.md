# Seabond AI 名词 & 面试题整理

> 基于 `seabond-api` 项目实际用到的 AI 架构整理。一个 OpenAI 兼容的交易类 AI agent：OpenRouter → Anthropic Claude，带原生 tool-calling、MCP、长期记忆、prompt caching、流式 agent loop。

---

## Part 1 — 核心名词表

### LLM 基础

| 名词 | 一句话解释 | 项目里怎么用 |
|---|---|---|
| **LLM（大语言模型）** | 在海量文本上训练、靠预测下一个 token 来生成文字的模型。 | 回复内容由 LLM 生成，默认 `anthropic/claude-sonnet-4.6` |
| **Token（词元）** | 模型读写的最小单位，约等于一个词的片段。 | 计费、quota、history budget 全部按 token 算 |
| **Context window（上下文窗口）** | 单次请求模型能关注到的最大 token 数。 | history 超过 30K token budget 就裁掉最老消息 |
| **Temperature（温度）** | 采样随机度，0 = 确定性，越高越发散。 | guard 分类器用 `temperature=0` 求确定性；chat 默认随上游 |
| **Max tokens（最大输出）** | 输出长度上限。 | 默认 4096 / 8192，可 per-request 覆盖 |
| **Prompt（提示词）** | 喂给模型的完整输入（system + 历史 + 用户）。 | 服务端拼装，前端不拼 |
| **System prompt（系统提示）** | 设定角色/规则的指令块，用户不可见。 | ~3KB，含 skill schema + 行为规则，按语言注入 |
| **Role（角色）** | 消息的发出方：`system` / `user` / `assistant` / `tool`。 | 客户端传的 `system` 会被服务端过滤掉 |
| **Finish reason / stop reason（结束原因）** | 生成为何停止：`stop` / `tool_calls` / `length`。 | 本项目扩展为 `stop` / `preview` / `tool_calls` / `blocked` |
| **Streaming / SSE（流式）** | Server-Sent Events，token 增量推送。 | `text/event-stream`，`data: <JSON>\n\n` 逐事件推 |
| **Multimodal（多模态）** | 模型能处理多于一种模态（文本 + 音频/图像）。 | 语音模式：音频上行 + 文本/音频下行 |

### Agent / Tool calling（工具调用）

| 名词 | 一句话解释 | 项目里怎么用 |
|---|---|---|
| **Tool / Function calling（工具/函数调用）** | 模型产出结构化调用（`名称 + JSON 参数`），由服务端执行。 | 走 native `tools[]` + `tool_calls[]` 协议，不再 emit `<action>` 文本 |
| **Agent loop（智能体循环）** | 循环：模型 → tool_call → 执行 → 回喂结果 → 模型，直到 `stop`。 | 服务端在 `/v1/chat/completions` 内部跑 streaming agent loop |
| **Tool executor（工具执行器）** | 真正去跑工具并返回结果的服务端组件。 | `AgentExecutor`，用用户身份调对应 service |
| **Tool result（工具结果）** | 序列化后的工具输出，以 `role:"tool"` 回喂模型。 | SSE 帧上限 2048 字节，截断只发前端；完整 payload 回喂模型 |
| **Step limit / max iterations（步数上限）** | 限制 agent loop 轮数，控成本控延迟。 | 默认 10 步；触顶 → `stopped:"tool_calls"` |
| **READ vs WRITE tool（读/写工具）** | 读 = 无副作用拉数据；写 = 改状态。 | READ 内联执行；WRITE 短路成 preview |
| **Preview tool / human-in-the-loop（预览工具 / 人在回路）** | 模型只提议一个变更动作，必须由人确认后才执行。 | 6 个写工具抛 `ErrAgentStop` 终结 loop，args 透传前端，用户签名才上链 |
| **JSON Schema（严格模式）** | 工具参数的类型契约，`additionalProperties:false`。 | 每个 ToolDef 带 strict schema |
| **MCP（Model Context Protocol）** | 把工具/资源暴露给任意 LLM 客户端的开放 JSON-RPC 标准。 | `POST /mcp`，长效 `mcp_` token 鉴权 |

### Prompt 工程 / 上下文管理

| 名词 | 一句话解释 | 项目里怎么用 |
|---|---|---|
| **Prompt caching（提示词缓存）** | 缓存稳定的 prompt 前缀，重复请求跳过重算。 | Anthropic `ephemeral` cache：base prompt 全用户共享，user context 单独 breakpoint |
| **Cache breakpoint（缓存断点）** | 标记可缓存前缀终点的边界。 | system 分两段：base（共享缓存）+ per-user context（只失效尾部） |
| **History truncation / token budget（历史裁剪 / token 预算）** | 丢弃最老的轮次以塞进预算。 | 超 30K token 裁最老消息，只影响喂 LLM 的 prompt，DB 全量保留 |
| **Long-term memory（长期记忆）** | 跨会话持久注入上下文的事实。 | `chat_memories`，per-user ≤200 条，TTL 60 天，`last_used_at` LRU |
| **Summarization（摘要）** | 用便宜模型把会话浓缩成可复用的笔记。 | summarizer 模型抽 0–2 条记忆 / 标题，异步 fire-and-forget |
| **SUPERSEDES / dedup（覆盖 / 去重）** | 新记忆把过期的旧记忆标记失效，防漂移。 | 写时去重，新记忆可标记旧记忆过期 |
| **Cheap/utility model（廉价/工具模型）** | 处理非用户可见子任务的小而快的模型。 | `skill-model`（如 gemini flash-lite）做摘要/分类/标题 |
| **Guardrail / pre-flight classifier（护栏 / 前置分类器）** | 主模型前的轻量检查，管成本/范围/安全。 | `ChatGuardService`：`ALLOW`/`INJECTION`(仅记日志)/`OFFTOPIC`(拦截) |
| **Prompt injection（提示词注入）** | 不可信输入劫持模型的指令。 | guard 分类 `INJECTION` 但只打日志、照常进 loop |
| **Hallucination（幻觉）** | 模型编造输入/工具里没有的"事实"。 | system prompt 要求只信工具结果、禁止编造/禁价格预测 |
| **Quota / rate limiting（配额 / 限流）** | per-user 用量上限以控成本。 | 默认 3M token/天/用户；断流/报错轮不计费 |
| **Idempotency（幂等）** | 同输入同效果，可安全重试。 | READ 工具无副作用、可重试；上游截断在首 token 前透明重试 2 次 |

### 部署 / 工程

| 名词 | 一句话解释 | 项目里怎么用 |
|---|---|---|
| **Gateway / router（网关，OpenRouter）** | 把多家 LLM 厂商统一到一个 API 后面的代理。 | OpenRouter → Anthropic；`go-sdk v0.4.1` |
| **OpenAI 兼容 API** | 请求/响应结构和 OpenAI `/chat/completions` 一致。 | 官方 `openai` SDK / LangChain 直接对接，换 baseURL + JWT |
| **Conversation / session（会话）** | 多轮线程，按用户隔离与标识。 | `chat_conversations`（UUID v7），侧边栏 ≤100 active |
| **Optimistic UI + real id（乐观 UI + 真实 id）** | 先立即渲染，等服务端 id 到了再对账。 | `X-User-Message-Id` header 提前回传真实消息 id |
| **Fire-and-forget（发后不管）** | 异步起任务、不阻塞响应。 | `afterRound`：标题生成 + 记忆汇总 + bump，单项失败只 WARN |
| **Backoff / retry（退避重试）** | 重试间隔逐步拉长。 | 上游断流退避 200ms→800ms |
| **Cache TTL by volatility（按波动率定缓存时长）** | 缓存寿命随数据变化速度调整。 | mark price 不缓存；新闻 30–60s；宏观 5–10min |

---

## Part 2 — 常见面试题（问答）

### A. LLM / Prompt 基础

**Q1. 什么是 token？和"词"有什么区别？**
token 是模型分词后的子词单位（英文平均 1 token ≈ 0.75 词）。计费、上下文窗口、配额都按 token 算，不按字符或词——所以成本随 token 数增长，不是随字数。

**Q2. `temperature` 控制什么？什么时候设 0？**
它缩放下一个 token 概率分布的 softmax。`temperature=0` ≈ 确定性（贪心）输出。分类/抽取/护栏检查这类要稳定可复现的决策就用 0（本项目前置 guard 用 `temperature=0`）；创意/对话用更高（0.7+）。

**Q3. 为什么这个 API 禁止客户端传 `role:"system"`？**
system prompt 编码了服务端控制的规则（工具 schema、安全约束、输出格式）。允许客户端注入 `system` 是 prompt-injection / 越狱向量，还会破坏 prompt 缓存。服务端会过滤掉并注入自己的。

**Q4. 上下文窗口是什么？历史超了怎么办？**
单次请求模型能关注的最大 token 数。这里历史按 30K token 预算裁剪，丢弃**最老**的消息——只裁喂给 LLM 的部分，DB 保留完整记录，UI 不受影响。

**Q5. 什么是幻觉？本系统怎么缓解？**
模型说出没有依据的"事实"。缓解：①system prompt 要求"用工具、别凭记忆答"；②所有行情数据来自确定性的工具结果，prompt 要求原样信任；③明令禁止价格预测（合规 + 抗幻觉）。

### B. Agent / Tool calling

**Q6. 解释一下 agent loop。**
注入上下文 → 模型产出 `tool_calls[]` → 服务端逐个用工具执行器执行 → 以 `role:"tool"` 回喂结果 → 模型继续 → 重复直到 `finish_reason==stop` 或触步数上限。它让模型现拉实时数据再推理，而不是靠陈旧权重作答。

**Q7. 为什么 WRITE 工具要短路成 "preview" 而不是直接执行？**
交易动真金白银且需要链上签名。服务端**绝不**自动执行写操作。模型调 `place_order` 等时，loop 抛 `ErrAgentStop`，返回 `{preview:true, tool, args}`，前端渲染确认卡，用户必须签名。这就是对不可逆动作的"人在回路"。

**Q8. 工具结果怎么返回？为什么 SSE 帧里要截断？**
结果序列化后以 tool 消息回喂模型。SSE 流里上限 2048 字节（`...(truncated)`），防止大体量 READ payload 把浏览器淹了——但**完整** payload 仍内部回喂模型。前端不要对大 READ 结果做 `JSON.parse`。

**Q9. 为什么 agent loop 要有步数上限？**
控成本控延迟、防工具调用死循环。触顶时该轮以 `stopped:"tool_calls"` 结束（罕见）。

**Q10. MCP 是什么？和 tool calling 什么关系？**
Model Context Protocol——把工具/资源暴露给任意 LLM 客户端的开放 JSON-RPC 标准。这里 `/mcp` 让外部 MCP 客户端（`mcp_` token 鉴权）调用内部 agent 用的同一套 READ 工具注册表，handler 一致性由测试保证。内部专用的组合原语会从公开 `tools/list` 过滤掉。

**Q11. 工具在 loop 中途报错怎么处理？**
把错误包成 `{error: "..."}` 回喂模型，让 loop 继续、模型自行恢复或解释——而不是让整个请求崩掉。

### C. Prompt caching / 成本

**Q12. 什么是 prompt caching？为什么把 system message 拆两段？**
Anthropic 缓存稳定的 prompt **前缀**（`ephemeral`），重复请求跳过重算。system message 拆成：①base prompt——所有用户一致，是共享缓存命中；②per-user context（记忆）——单独缓存断点，只让用户专属尾部失效。在不丢个性化的前提下最大化缓存命中率。

**Q13. 为什么不能直接缓存整个 `/v1/chat/completions` 响应？**
输入每轮都变（历史增长 + 记忆漂移），输出非确定性，且回包带 `tool_calls` 会驱动真实交易 + 必须真跑的副作用。缓存下沉到更细的层：prompt 前缀缓存 + 按 `hash(name,args)` 的 per-tool Redis 缓存，TTL 随数据波动率调。

**Q14. 怎么降低 agent 系统的 LLM 成本？（开放）**
prompt 前缀缓存；历史按 token 预算裁剪；廉价子任务（摘要/分类/标题）路由到小"skill"模型；确定性工具输出按波动率缓存；中断/报错轮不计配额；给 loop 设步数上限。

### D. 记忆 / 上下文

**Q15. 长期记忆 vs 会话历史，区别是什么？**
历史 = 当前会话的逐字轮次（受 token 预算约束）。长期记忆 = 便宜模型从过往轮次蒸馏出的可复用事实，跨**所有**会话注入（per-user ≤200，TTL 60 天，靠 `last_used_at` 做 LRU）。

**Q16. 怎么防止记忆过期/自相矛盾？**
`SUPERSEDES` 机制：新记忆可把旧的标记失效，写时做去重，TTL + LRU 淘汰冷条目、注入时让热条目续命。

### E. 安全 / 可靠性

**Q17. guard 分类器和安全边界的区别？**
guard 是**成本/范围**过滤器，不是安全边界。`OFFTOPIC`（写代码/写作文/白嫖 LLM）拦截；`INJECTION` 只记日志（祈使语气的下单会误触）仍进 loop；确定性 fast-path 让明确交易意图跳过分类器。真正的安全 = WRITE 工具绝不自动执行 + 用户签名模型。

**Q18. 上游流在生成中途被截断，系统该怎么表现？**
若在任何客户端可见输出**之前**截断 → 透明重试该轮（≤2 次，退避 200→800ms），客户端只多等 ~1s。若在已吐部分输出**之后**截断 → 推 `error` 事件，**不**伪造 `done`，不落库不计配额。绝不把截断流当成功。

**Q19. 客户端中途断开，落库什么、计费什么？**
已生成的半截 assistant 文本带本地化中断标记落库（刷新仍在，摘要器据此知道没说完），但 `totalTokens=0` → **不扣配额**。纯 tool-only 无文本的断开跳过空 assistant 行。

**Q20. 为什么 user 消息要在 agent loop 启动**之前**落库？**
这样无论后续 loop 报错、断流还是 tool-only 轮，用户问的话都不丢。落库分两段：user 消息提前写，assistant 行在 `afterRound`（clean done 后）写——避免"有会话行、零消息"的空会话。

### F. 系统设计（开放题）

**Q21. 设计一个 OpenAI 兼容、底层 Claude 带工具的 chat API。（要点）**
OpenAI 形态的 `/v1/chat/completions` → 过滤客户端 `system` → 拼装 prompt（缓存 base + per-user 记忆 + token 预算内的历史）→ guard 前置 → 流式 agent loop（原生工具调用，READ 内联 / WRITE preview）→ SSE 事件 → 两段式落库 → 异步 `afterRound`（标题 + 记忆摘要）→ 仅 clean 轮计 token 配额。

**Q22. 怎么把工具调用流式做得对前端可观测？**
通过 SSE 推有类型的生命周期事件：`tool_call_start` → `tool_call_args_delta`（增量 JSON）→ `tool_call_done` → `tool_call_executing` → `tool_call_result` → `text_delta` → `done {stopped,...}` → `[DONE]`。前端只认 `done.stopped`，绝不靠 `is_error` 判定。

**Q23. 厂商抽象：为什么走 OpenRouter 而不直连 Anthropic？**
网关把多厂商统一到一个 OpenAI 兼容 API 后面：支持 per-request 模型覆盖（如按请求钉到音频模型而不影响其它请求）、厂商故障转移；客户端只需换 `baseURL` + key 就能直接用标准 OpenAI SDK / LangChain。

---

## Part 3 — 高频英文表达（面试口语）

> 名词表/答案已全中文；这一节保留英文表达供英文面试时背诵，括号内为中文。

- "We run a **streaming agent loop** server-side — the model emits native `tool_calls`, we execute them, feed the results back, and it continues until it stops."（服务端跑流式 agent loop：模型出原生 tool_calls，我们执行、回喂、继续直到停止。）
- "READ tools are executed inline; **WRITE tools are short-circuited into a preview** that the user has to confirm and sign — the server never auto-executes a trade."（读工具内联执行；写工具短路成 preview，用户确认并签名，服务端绝不自动下单。）
- "We split the system prompt into a **shared cached prefix** and a **per-user cache breakpoint** to maximize the **prompt cache hit rate** without losing personalization."（system prompt 拆成共享缓存前缀 + per-user 缓存断点，最大化缓存命中又不丢个性化。）
- "History is trimmed to a **token budget** by dropping the oldest turns — only the prompt fed to the model is cut; the full transcript stays in the DB."（历史按 token 预算裁，丢最老轮次；只裁喂模型的部分，DB 留全量。）
- "Cheap subtasks like summarization, classification, and title generation are **routed to a small utility model**, not the main agent model."（摘要、分类、标题这类廉价子任务路由到小工具模型，不走主 agent 模型。）
- "An aborted or errored round **doesn't burn quota** — we only charge on a clean completion."（中断或报错的轮次不扣配额，只在干净完成时计费。）
- "The pre-flight guard is a **cost-and-scope filter, not a security boundary**; the real safety guarantee is human-in-the-loop signing for any mutating action."（前置 guard 是成本/范围过滤器、不是安全边界；真正的安全保障是任何变更动作都人在回路签名。）
- "On upstream truncation we **transparently retry before any client-visible output**, but surface an `error` once partial output has been sent — we never fake a success."（上游截断时，在客户端可见输出前透明重试；已吐部分输出后则推 error，绝不伪造成功。）

---

_整理自 seabond-api 实际代码与 `docs/api_ai_chat.md`。日期 2026-05-19。_
