## Goroutine and  Channel
```
Goroutine（协程）: Goroutine 是一个轻量级线程，它允许在单个 Go 程序中并发运行数十万个 Goroutine
Channel 是 Go 的 通信机制，用于 goroutine 之间安全地传递数据。
Go 的并发哲学：不要通过共享内存来通信，而通过通信来共享内存
常见模式：
工作池（Worker Pool）：主 goroutine 生成任务，发送到 channel，worker goroutine 从 channel 中取任务并处理。

Goroutine：Goroutine is a lightweight thread，it allows hundreds of thousands of goroutines to run concurrently in a single Go program
Channel： Channel is Go’s communication mechanism, used for safely passing data between goroutines.
Go’s concurrency philosophy: Do not communicate by sharing memory; instead, share memory by communicating.

Typical example (bin):
Worker Pool: the main goroutine generates tasks and sends them to a channel, 
while worker goroutines retrieve tasks from the channel and process them.

```