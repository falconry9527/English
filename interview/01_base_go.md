## Goroutine and  Channel
```
Goroutine（协程）: Goroutine 是一个轻量级线程，它允许在单个 Go 程序中并发运行数十万个 Goroutine
Channel 是 Go 的 通信机制，用于 goroutine 之间安全地传递数据。
Go 的并发哲学：不要通过共享内存来通信，而通过通信来共享内存
常见模式：
工作池（Worker Pool）：主 goroutine 生成任务，发送到 channel，worker goroutine 从 channel 中取任务并处理。

Goroutine：Goroutine is a lightweight thread，it allows thousands of goroutines to run concurrently in a single Go program
Channel： Channel is Go’s communication mechanism, used for safely passing data between goroutines.
Go’s concurrency philosophy: Do not communicate by sharing memory; instead, share memory by communicating.

Typical example (bin):
Worker Pool: the main goroutine generates tasks and sends them to a channel, 
while worker goroutines get tasks from the channel and process them.

```

``` 其他常见面试题
Q1:数组 (array) 和切片 (slice) 的区别？
数组是固定长度，值类型；切片是动态长度，引用类型，底层指向数组。

Q2:context 包的作用？
管理 goroutine 生命周期，支持超时、取消、传递元数据。

Q3:Go 的内存逃逸分析是什么？
编译器决定变量在栈上还是堆上分配。

Q4:Go 的垃圾回收机制？
三色标记清除（Mark-Sweep），并发 GC，减少 STW 时间。

Q5:Go 在分布式系统中如何保证服务稳定性？
超时控制 (context)、限流、熔断、重试机制。

Q6:gin 框架的原理？
基于 radix tree 的路由，支持中间件链式调用。
Radix Tree 是一种基于 字符串前缀压缩 的搜索树。Radix Tree 会把公共前缀合并，大幅减少空间开销。

Q7:Go 的并发安全问题如何解决？
使用 channel 通信，或者使用 sync 包（如 Mutex、RWMutex、WaitGroup）。

Q1: What’s the difference between arrays and slices in Go?
Arrays have a fixed length and are value types.
Slices have a dynamic length, are reference types, and internally point to an array.

Q2: What’s the purpose of the context package?
It manages the lifecycle of goroutines, supporting timeout, cancellation, and passing metadata.

Q3:What is Go’s escape analysis?
The compiler decides whether a variable should be allocated on the stack or the heap.

Q4:What’s Go’s garbage collection mechanism?
Go uses a concurrent tri-color mark-sweep (Mark-Sweep) GC algorithm, which minimizes stop-the-world (STW) pauses.

Q5:How does Go ensure stability in distributed systems?
By using timeout control (context), rate limiting, circuit breakers, and retry mechanisms.

Q6: What’s the principle behind the Gin framework?
Gin is based on a radix tree for routing and supports middleware chaining.
A Radix Tree is a prefix-compressed trie that merges common prefixes to reduce memory usage.

Q7: How does Go handle concurrency safety?
By using channels for communication, or synchronization primitives from the sync package such as Mutex, RWMutex, and WaitGroup.

```