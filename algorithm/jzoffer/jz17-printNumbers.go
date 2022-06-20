package main

import (
	"fmt"
	"math"
)

/**
 * [JZ17-简单] 打印从1到最大的n位数
 *
 * @param n int整型 最大位数
 * @return int整型一维数组
 */
func printNumbers(n int) []int {
	pow := math.Pow10(n)
	count := int(pow) - 1
	arr := make([]int, count)
	for i := 1; i <= count; i++ {
		arr[i-1] = i
	}
	return arr
}

func main() {
	ret := printNumbers(1)
	fmt.Println(ret)
}
