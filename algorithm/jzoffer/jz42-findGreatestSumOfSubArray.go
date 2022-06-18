package main

import (
	"fmt"
	"math"
)

/**
 * [JZ44-简单] 数字序列中某一位的数字
 *
 * @param array int 一维数组
 * @return int
 */
func findGreatestSumOfSubArray(array []int) int {
	length := len(array)
	if 1 == length {
		return array[0]
	}

	dp := make([]int, length)
	dp[0] = array[0]
	max := array[0]
	for i := 1; i < length; i++ {
		dp[i] = int(math.Max(float64(dp[i-1]+array[i]), float64(array[i])))
		max = int(math.Max(float64(dp[i]), float64(max)))
	}
	return max
}

func main() {
	arr := []int{1, -2, 3, 10, -4, 7, 2, -5}
	ret := findGreatestSumOfSubArray(arr)
	fmt.Println(ret)
}
