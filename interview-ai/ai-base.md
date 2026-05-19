
# 常用名词
```
AIGC = AI-Generated Content（人工智能生成内容）-> 内容生成式AI

Agent: （智能体） 具有思考决策，记忆，调用工具 能力的系统。
一个 Agent 通常由四部分组成
LLM（大脑）:	推理、决策下一步做什么	Claude（OpenRouter）
Tools（工具/手脚）	: 暴露给LLM的外部能力（skill）：查数据 READ 工具 + 下单 WRITE preview 工具、MCP（Chatbot 只会聊）
Memory（记忆）: 	短期=对话历史，长期=跨会话事实	chat_messages 历史 + chat_memories 长期记忆
Loop（循环）:  一次请求，自己调整，多次调用大模型	服务端 streaming agent loop

MCP(模型上下文协议):主要用来给 LLM 暴露能力，比如，抓取网页，获取持仓等

OpenRouter : 一个 LLM 聚合路由(router),把 chatgpt、claude、gemini、deepseek 等几百个模型统一到一起

prompt(系统提示词) : prompt 是 用户消息  + 记忆 + 历史 + system + 工具的完整上下文整体。
用户消息是 prompt 的输入之一;

```

# 怎么节省token
```
history 预算裁剪(省历史税) : 
多模型分流(降单价) : chat 的大模型llm，memroy模型，标题生成，新闻分析,前置Guard 分类器 。
Guard 前置拦截(省整次调用) : 只允许交易相关的会话，避免成为 其他些代码，写小说的工具 。

prompt caching(省重复输入) : 大行情分析，缓存1分钟，减少ai 的多次调用

```

# 其他没有用到的 
```
ARG : 检索增强生成,回答前先去内部知识库检索相关内容 : 非结构化文档里,且文档量大/更新快/私有


```
