package main

import (
	"fmt"
	"math"
)

/**
 * [JZ63-简单] 买卖股票的最好时机(一)
 *
 * @param array int 一维数组
 * @return int
 */
func maxProfit(prices []int) int {
	n := len(prices)
	if 0 == n {
		return 0
	}

	min := prices[0]

	ret := 0
	for i := 1; i < n; i++ {
		min = int(math.Min(float64(min), float64(prices[i])))
		ret = int(math.Max(float64(ret), float64(prices[i]-min)))
	}
	return ret
}

func main() {
	arr := []int{8, 9, 2, 5, 4, 7, 1}
	ret := maxProfit(arr)
	fmt.Println(ret)
}
