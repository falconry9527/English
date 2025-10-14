## go 语言的数据类型
```
一、值类型（Value Types）
整型（Integer）：int, int8, int16, int32, int64, 
以及对应的无符号类型 uint8, uint16, uint32, uint64, uintptr。
浮点型（Float）：float32, float64。
复数类型（Complex）：complex64, complex128。
布尔型（Boolean）：bool，取值为 true 或 false。
字符串类型（String）：string，表示不可变的字节序列。
数组类型（Array）：[N]T，固定长度的同类型元素序列，属于值类型，拷贝时会复制整个数组内容。

二、引用类型（Reference Types）
切片（Slice）：[]T，动态数组，底层引用一个数组。
映射（Map）：map[K]V，键值对存储结构。
通道（Channel）：chan T，用于 goroutine 之间的通信。
接口（Interface）：interface{}，可以存储任意类型的值，实现多态。
函数类型（Func）：func()，函数本身也可以作为变量、参数或返回值。

三、其他类型
指针（Pointer）：*T，保存变量的内存地址，通过指针可以间接访问或修改数据。
结构体（Struct）：struct {}，自定义的复合数据类型，用于组合多个字段。
自定义类型（Type Alias / Type Definition）：type MyInt struct，基于已有类型定义新类型或类型别名。

```

## slice
```
一. 定义和初始化: slice 是一个 封装类型，封装，指针（指向底层数组）,len(当前长度),cap (当前容量)
type slice struct {
    array unsafe.Pointer // 指向底层数组
    len   int            // 当前长度
    cap   int            // 当前容量
}
不指定初始长度的时候，初始没有分配桶，一次插入元素时才真正分配内存（懒加载）。
指定初始长度的时候，就会根据容量分配足够的桶
数据量大于初始容量也不会报错，因为会自动扩容

二. 扩容
触发条件
s = append(s, x)
当容量小于 1024 时: 每次扩容 翻倍（×2）
当容量 ≥ 1024 时: 每次扩容大约增长 1.25 倍（+25%）
```

## map 
```
一. 定义和初始化
Go 的 map 是基于 哈希表（hash table）和 桶（bucket） 实现的 ,每个桶能存 最多 8 个键值对。
Go 的 map 
如果不指定初始容量，初始没有分配桶。第一次插入元素时才真正分配内存（懒加载）。
如果不指定初始容量，就会根据容量分配足够的桶 。
数据量大于初始容量也不会报错，因为会自动扩容

二. 扩容
1.等量扩容 : 数据分布不均匀，保持桶数量不变，只做 rehash
2.加倍扩容 : 装载因子过大(> 6.5), 即平均每个桶中元素太多,会触发扩容
```

## slice,map 创建的时候，应该指定长度吗
```
如果知道长度，建议 指定，减少多次扩容的消耗
```

## make
```
切片: 
nums : = make([]int, 3) // slice 必须指定长度
map: 
m : = make(map[string]int, 10) // map可以不指定长度，初始化长度为0
channel: 
chan1 : = make(chan int32) // channel 不指定长度，就是无缓冲区channel
```

## interface
```
参考代码： 02_interface
一. 定义 和 数据结构
interface 是一种 类型抽象

interface 的底层是 两部分组成
类型信息（type）
数据指针（data）
s : = i.(string) // 断言 i 是 string
v : = i.(type)

二.分类和作用: 
通用容器: 空接口 interface{} 不包含任何方法 : 可以存储 任意类型的值
多态与抽象: 带方法的接口: Person 实现了 Speaker 接口，隐式实现，不需要显式声明
```


## go 的并发机制-csp模型
```
go 的并发机制是基于CSP（Communicating Sequential Processes）通信顺序进程
核心理念：通过通信来共享内存,实现并发

Go 的CSP模型主要由 Goroutine + Channel组成：
goroutine : 轻量级线程 
channel : 是 Go 的 通信机制，用于 goroutine 之间安全地传递数据。

典型案例: 
工作池（Worker Pool）: 主 goroutine 生成任务，发送到 channel，worker goroutine 从 channel 中取任务并处理。
```

## channel
```
Channel 是 Go 的 通信机制，用于 goroutine 之间安全地传递数据。
无缓冲 channel → 发送和接收必须同时完成（同步通信），写入的信息的时候: 一直阻塞直到有接收
带缓冲 channel → 异步通信，缓冲区满时发送阻塞

select 可以同时等待多个 channel,类似于switch
```


## Goroutine
```
Goroutine（协程）: Goroutine 是一个轻量级线程,协程占用的内存空间极小，允许在单个Go程序中并发运行数十万个 Goroutine 。

协程是不为操作系统所感知的，它由编程语言层面实现，上下文切换(协程切换)不需要经过内核态 

什么是 内核态 呢，比如 ，java 的线程 
Java 的线程 是由操作系统内核所感知的，它由操作系统层面实现，上下文切换(线程切换)需要经过内核态。
CPU 从线程 A 切到线程 B 时 ，保存线程 A 的寄存器、栈等上下文，恢复线程 B 的寄存器、栈等上下文
java21（23年）以后，也引入了 虚拟线程，其实就是协程

Goroutine 的切换不是 内核态 的，具体是  GPM 模型：
```


## GPM 调度模型
```
协程-逻辑处理器-线程
G（Goroutine）: 协程，表示一个轻量级线程，用户代码逻辑运行单元
P（Processor）: 逻辑处理器，管理调度资源，负责分配 G 到 M 上运行。
M（Machine）: 操作系统线程，真正执行 G 的操作系统线程。

具体流程 :
G 创建 ：
go func() 会创建一个 G 并放入某个 P 的 本地队列 （local run queue） 或 全局队列（global run queue）。

P 分配 ：
P 维护 一个本地队列 ，优先执行本地队列中的 G。
如果 本地队列空，则从 global queue 或其他 P 的队列窃取 G（work stealing）。

M 执行 ：
每个 M 绑定一个 P。
M 从 P 获取 G 并执行。
如果 G 阻塞或调用系统调用（如网络 IO），M 会释放 P，去执行其他 P 上的 G
（这个过程会存在内核切换，但是存在的可能性很低，几乎为0）。
被阻塞的 G 会被 P 放入等待队列（wait queue），等待唤醒。
```

## channel : 对已经关闭的的chan进行读写
```
channel 关闭后，不能再写（send），但还能读（receive）剩余的数据。
对已关闭的channel
写: 会直接报错: (panic: send on closed channel)
读: 
1. 正常读取剩余的数据
2. 不会报错，返回值返回读取失败 false（读取的时候有两个返回值: 数据，是否读取成功）
    v, ok : = <-ch （OK 为false）
```

## 内存逃逸
```
定义: 
Go 程序中的 局部变量 通常分配在栈（stack） 上 ,
但是，如果局部变量 在 函数返回后 仍然被使用， 就会被分配在堆（heap） 上，由垃圾回收器（GC）管理。
这个从「栈 → 堆」的过程，就叫做 内存逃逸（memory escape）。

危害: 
逃逸频繁 → GC 压力大，内存占用越来越高，最终可能 OOM 或者性能降级。

案例: （排查代码有没有一些常见的内存逃逸错误）
// 1. 返回局部变量的指针
func foo() *int {
	x : = 10
	return &x // x 逃逸到堆上
}
// 2. 将变量存入 interface
借口的data是指针类型
func bar() interface{} {
	s1 : = "hello"
	return s1 // s 逃逸
}
// 3. 闭包引用外部变量（全局变量）
func main2() {
	var s2 string
	f : = func() {
		fmt.Println(s2) // s 可能逃逸
	}
	f()
}
// 4. slice 或 map 导致内部结构被堆分配
func test() {
	a : = make([]int, 1000) // 可能在堆上分配
	println(a)
}

查看是否内存逃逸命令: 
go build -gcflags=-m 05_memory_escape.go

GC 日志
GODEBUG=gctrace=1 ./your_program

```

## go 内存泄漏排查
```
一.定义:
程序 长期引用某些对象,没有释放（通过全局变量、goroutine、Channel、切片、map 等），
导致 GC 无法回收，内存占用越来越高，最终可能 OOM 或者性能降级。

二.案例 （排查代码有没有一些常见的内存逃逸错误）
1. Goroutine 泄漏 :  
Goroutine 没有及时关闭，导致数量越来越多
代码中，注意关闭 Goroutine
a. 新增一个channel，发送 done 信号，让 worker Goroutine 安全退出；
b. 使用 context.Context 控制生命周期，关闭 Goroutine；

2. Channel 阻塞/未关闭
使用无缓冲区的channel 或者 Channel 未关闭
a. 使用有缓冲区的channel，避免阻塞
b. defer 及时关闭  channel 

3. 切片保留底层数组
go big := make([]byte,10_000_000); small := big[:10]; big=nil
拷贝需要的数据到新切片：smallCopy := append([]byte(nil), small...)

4. Map / 缓存无限增长
go cache[key]=value，长时间不删除
定期清除 不用的 key

5. 未关闭的资源 : 文件，数据库链接，imer/Ticker
f, _ := os.Open("file") 后未 defer f.Close()
及时 Close()；使用 defer(dɪˈfɜː(r)) 确保释放

6. 闭包捕获大对象
go func() { fmt.Println(hugeArray) }()
避免捕获大对象；只传必要数据；或在闭包外拷贝小对象

三. 解决方法： 启用 pprof 分析
a.查找长期阻塞的 goroutine,channel
b.查看哪些对象、函数或类型持续增长。

1. 在程序中导入并启动 pprof：
import _ "net/http/pprof"
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
2. 收集 heap profile：
go tool pprof http://localhost:6060/debug/pprof/heap > heap1.pb
# 等待一段时间或高负载运行后
go tool pprof http://localhost:6060/debug/pprof/heap > heap2.pb

3. 分析差异：
go tool pprof -diff_base=heap1.pb heap2.pb
查看哪些对象、函数或类型持续增长。

4. 查看 goroutine 堆栈：
查找长期阻塞的 goroutine（channel 阻塞、select 阻塞、等待信号）。

```

## go 的垃圾回收（gc）
```
定义:
Golang 的垃圾回收是 自动管理内存 的机制，它会自动检测 堆上 不再使用的对象并释放它们的内存。

过程：
并发标记-清除（concurrent mark-and-sweep） 型垃圾回收器。
1. 标记（Mark） : 标记所有可达的对象(黑色)。
2. 清除（Sweep） : 清理未被标记的对象。 
3. 触发条件 : 堆大小增长 100%

GC 的调度方式 : 
三色标记法 + 并发回收
三色标记法：
白色：初始化颜色
灰色：正在扫描的对象
黑色：已扫描且可达的对象
最后回首的是： 白色的对象

并发回收：
GC 标记阶段大部分与应用 goroutine 并发运行

```


## go 的 生命周期
```
Created ---> Ready ---> Running ---> Blocked(阻塞) ---> Terminated（结束）
可能的阻塞操作：
channel 阻塞：等待发送或接收数据
锁阻塞：sync.Mutex 或 sync.WaitGroup 等
时间等待：time.Sleep、time.After
```

## go 的 上下文
```
context： 管理 goroutines 的生命周期
例如 : 
取消 goroutine → 防止泄漏
超时 / deadline → 控制操作时长
传递元数据 → 轻量级共享请求信息

context 传递元数据 : 跨 goroutine 传递 请求相关信息，例如：
用户 ID、Session ID、Trace ID（分布式跟踪）
权限信息（Read/Write 权限标记）
```


## sync 包
```
Mutex（互斥锁）
Mutex 是独占锁，在锁持有期间，无论是读还是写，都会被阻塞。
适合 读写比例接近或写操作较频繁 的场景。

RWMutex（读写锁） ： 适合读操作教为频繁的场景
读锁（RLock）：
共享锁，多个 goroutine 可以共享一个读锁，并发读取不会互相阻塞。
但是 在持有写锁期间，新的读操作会被阻塞。
读锁的本质目的不是禁止并发读，而是防止在写入时发生读取，保证数据一致性。
写锁（Lock）：
独占锁，在写锁持有期间，所有读写操作都会被阻塞。

WaitGroup : Go 提供的一个 等待一组 goroutine 完成 的机制。
wg.Add(1) // 启动一个 goroutine，加1
defer wg.Done() // 结束时减1
wg.Wait() // 等待所有 goroutine 完成
```

## Printf
```
Printf —— 打印到控制台（标准输出）
Sprintf —— 格式化为字符串（不会打印）
Fprintf —— 写入到文件
```

## new 与 make 的区别（和指针有关）
```
动作，返回，适用类型
new：分配一块内存但不初始化，返回指向零值的指针，适用于值类型（如 int、float、struct、数组等）
make：创建并初始化对象，返回对象本身，适用于引用类型（slice、map、channel）.
```

