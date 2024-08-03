package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer func() {
		fmt.Println("done.")
	}()
	fmt.Println("CPU: ", runtime.NumCPU())
	//fmt.Println(int(math.Pow(2, 32)) / (1024 * 1024 * 1024))
}
