package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter1 int = 0

func add4(a, b int, lock *sync.RWMutex) {
	c := a + b
	lock.Lock()
	counter1++
	fmt.Printf("%d: %d + %d = %d\n", counter1, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go add4(1, i, lock)
	}

	for {
		// 读锁锁定会阻塞其他写锁，但不会阻塞读锁
		lock.RLock()
		c := counter1
		lock.RUnlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
