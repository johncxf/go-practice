package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(o *sync.Once) {
	fmt.Println("Start:")
	o.Do(func() {
		fmt.Println("Do Something...")
	})
	fmt.Println("Finished.")
}

func main() {
	o := &sync.Once{}
	go doSomething(o)
	go doSomething(o)
	time.Sleep(time.Second * 1)
}
