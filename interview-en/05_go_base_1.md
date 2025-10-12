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
切片：
nums := make([]int, 3) // slice 必须指定长度
map：
m := make(map[string]int, 10) // map可以不指定长度，初始化长度为0
channel：
chan1 := make(chan int32) // channel 不指定长度，就是无缓冲区channel
```

## interface
```
一. 定义 和 数据结构
interface 是一种 类型抽象

interface 的底层是 两部分组成
类型信息（type）
数据指针（data）
s := i.(string) // 断言 i 是 string
v := i.(type)

二.分类和作用:
通用容器：空接口 interface{} 不包含任何方法 ： 可以存储 任意类型的值
多态与抽象： 带方法的接口：Person 实现了 Speaker 接口，隐式实现，不需要显式声明
```

## Goroutine and  Channel
```
Goroutine（协程）: Goroutine 是一个轻量级线程,协程占用的内存空间极小，允许在单个Go程序中并发运行数十万个 Goroutine 。

协程是不为操作系统所感知的，它由编程语言层面实现，上下文切换(协程切换)不需要经过内核态 。
Java 的线程是由操作系统内核所感知的，它由操作系统层面实现，上下文切换(线程切换)需要经过内核态。
CPU 从线程 A 切到线程 B 时 ，保存线程 A 的寄存器、栈等上下文，恢复线程 B 的寄存器、栈等上下文

Channel 是 Go 的 通信机制，用于 goroutine 之间安全地传递数据。
Go 的并发哲学：不要通过共享内存来通信，而通过通信来共享内存
常见模式：
工作池（Worker Pool）：主 goroutine 生成任务，发送到 channel，worker goroutine 从 channel 中取任务并处理。

```

## channel 
```
无缓冲 channel :写入的信息的时候: 一直阻塞直到有接收
有缓冲 channel :写入的信息的时候 : 如果缓冲区慢，发生阻塞
```

## channel ： 对已经关闭的的chan进行读写
```
channel 关闭后，不能再写（send），但还能读（receive）剩余的数据。
对已关闭的channel
写：会直接报错: (panic: send on closed channel)
读：
1. 正常读取剩余的数据
2. 不会报错，返回值返回读取失败（读取的时候有两个返回值： 数据，是否读取成功）
    v, ok := <-ch （OK 为false）
```


## 内存逃逸
```
定义：
Go 程序中的局部变量通常分配在 栈（stack） 上，因为栈上的内存分配和释放速度非常快（函数返回时自动回收）。
但是，如果局部变量在函数返回后仍然可能被使用，它就不能放在栈上， 而是放在堆（heap） 上，由垃圾回收器（GC）管理。
这个从「栈 → 堆」的过程，就叫做 内存逃逸（memory escape）。
危害：
逃逸频繁 → GC 压力大：频繁堆分配会增加垃圾回收次数，从而降低程序吞吐量。
案例：
// 1. 返回局部变量的指针
func foo() *int {
	x := 10
	return &x // x 逃逸到堆上
}
// 2. 将变量存入 interface
func bar() interface{} {
	s1 := "hello"
	return s1 // s 逃逸
}
// 3. 闭包引用外部变量
func main2() {
	var s2 string
	f := func() {
		fmt.Println(s2) // s 可能逃逸
	}
	f()
}
// 4. slice 或 map 导致内部结构被堆分配
func test() {
	a := make([]int, 1000) // 可能在堆上分配
	println(a)
}

查看是否内存逃逸命令：
go build -gcflags=-m 05_memory_escape.go

```

## GPM调度模型
```
G（Goroutine）：Go 的协程，表示一个轻量级线程，它对应一个结构体g，结构体里保存了goroutine的堆栈信息
M（Machine）：操作系统线程，实际执行 goroutine 的物理载体。
P（Processor）：逻辑处理器，负责管理和调度 G 到 M 的映射。

```

