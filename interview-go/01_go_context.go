package main

import (
	"context"
	"fmt"
	"time"
)

// 模拟一个耗时操作
func longTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // 超时
		fmt.Println("Task completed")
	case <-ctx.Done(): // 取消
		fmt.Println("Task cancelled:", ctx.Err())
		return
	}
}

func main() {
	// 设置超时时间为 3 秒
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // 确保资源释放

	go longTask(ctx)

	// 等待 goroutine 结束
	time.Sleep(6 * time.Second)
	fmt.Println("Main exits")
}
