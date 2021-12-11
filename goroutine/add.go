// 协程实现加法示例
package main

import (
	"fmt"
	"time"
)

// 两数相加
func add(a, b int) {
	var c = a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

// main 函数是一个主协程
func main() {
	for i := 0; i < 10; i++ {
		// 协程调用
		go add(1, i)
	}
	// 让主协程等待 1s 后退出
	time.Sleep(1e9)
}
