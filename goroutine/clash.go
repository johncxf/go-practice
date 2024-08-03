package main

import (
	"fmt"
	"time"
)

var num = 0

func add() {
	for i := 0; i < 10000; i++ {
		num++
	}
}

func main() {
	// 演示未上锁情况下协程冲突情况
	go add()
	go add()
	time.Sleep(3 * time.Second)
	// 不冲突情况应该输出 20000
	fmt.Printf("num: %d", num)
}
