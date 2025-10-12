package main

import (
	"context"
	"fmt"
	"time"
)

// worker 接收任务，监听 done 或 ctx 退出信号
func worker(ctx context.Context, id int, jobs <-chan int, done <-chan struct{}) {
	for {
		select {
		case job := <-jobs:
			// 模拟任务处理
			fmt.Printf("Worker %d started job %d\n", id, job)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d finished job %d\n", id, job)

		case <-done:
			// 收到 done 信号，安全退出
			fmt.Printf("Worker %d received done signal, exiting\n", id)
			return

		case <-ctx.Done():
			// context 被取消，也安全退出
			fmt.Printf("Worker %d context cancelled, exiting\n", id)
			return
		}
	}
}

func main() {
	// 1️⃣ 使用缓冲 channel，防止发送阻塞
	jobs := make(chan int, 5)

	// 2️⃣ done channel，用于广播退出信号
	done := make(chan struct{})

	// 3️⃣ context，用于统一控制 goroutine 生命周期
	ctx, cancel := context.WithCancel(context.Background())

	// 启动 worker
	for i := 1; i <= 3; i++ {
		go worker(ctx, i, jobs, done)
	}

	// 发送任务
	for j := 1; j <= 5; j++ {
		fmt.Println("Sending job", j)
		jobs <- j
	}

	// 模拟任务处理一段时间
	time.Sleep(6 * time.Second)

	// 4️⃣ 方法一：发送 done 信号，让 worker 安全退出
	fmt.Println("Sending done signal to workers")
	close(done)

	// 5️⃣ 方法二：也可以使用 context 统一取消
	cancel() // 如果取消 context，也能触发 ctx.Done() 让 worker 退出

	// 再等待一段时间，确保 worker 完全退出
	time.Sleep(time.Second)
	fmt.Println("Main exits")
}
