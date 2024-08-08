package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("done.")
	}()
	//fmt.Println("CPU: ", runtime.NumCPU())
	//fmt.Println(int(math.Pow(2, 32)) / (1024 * 1024 * 1024))
	a := 0.1
	b := 0.2
	c := a + b
	fmt.Println(c)

	//arr := [26]int{}
	arr := make([]int, 26)
	arr[25] = 1
	fmt.Println(arr)
}
