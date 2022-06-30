package main

import "fmt"

/**
 * [JZ62-中等] 孩子们的游戏(圆圈中最后剩下的数)
 *
 * @param n int 整型
 * @param m in t整型
 * @return int 整型
 */
func lastRemaining(n int, m int) int {
	x := 0
	for i := 2; i <= n; i++ {
		x = (m + x) % i
	}
	return x
}

func main() {
	ret := lastRemaining(5, 3)
	fmt.Println(ret)
}
