package main

import (
	"fmt"
	"os"
)

func main() {
	// 示例数据
	name := "Alice"
	age := 25

	// 1️⃣ Printf —— 打印到控制台（标准输出）
	fmt.Printf("[Printf] Name: %s, Age: %d\n", name, age)

	// 2️⃣ Sprintf —— 格式化为字符串（不会打印）
	result := fmt.Sprintf("[Sprintf] Name: %s, Age: %d", name, age)
	fmt.Println(result) // 打印出格式化后的字符串

	// 3️⃣ Fprintf —— 写入到文件
	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "[Fprintf] Name: %s, Age: %d\n", name, age)

	fmt.Println(" 内容已写入 output.txt 文件")
}
