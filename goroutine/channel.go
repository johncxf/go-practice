// 通过 channel 进行消息传递
package main

import (
	"fmt"
	"time"
)

func add2(a, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	// 发送数据到 channel
	ch <- 1
}

func main() {
	start := time.Now()
	// 定义通道类型切片
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go add2(1, i, chs[i])
	}
	for _, ch := range chs {
		// 接收数据
		<-ch
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
