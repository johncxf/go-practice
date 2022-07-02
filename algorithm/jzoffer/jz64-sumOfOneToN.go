package main

import "fmt"

/**
 * [JZ64-中等] 求1+2+3+...+n
 *
 * @param n int 整型
 * @return int 整型
 */
func sumOfOneToN(n int) int {
	return n * (n + 1) / 2
}

func main() {
	fmt.Println(sumOfOneToN(100))
}
