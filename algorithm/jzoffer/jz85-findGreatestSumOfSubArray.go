package main

import "fmt"

/**
 * [JZ42-简单] 连续子数组的最大和
 *
 * @param array int 一维数组
 * @return int
 */
func findGreatestSumOfSubArr(array []int) []int {
	length := len(array)
	if 1 == length {
		return array
	}

	dp := make([]int, length)
	dp[0] = array[0]
	max := array[0]

	// 滑动左右区间值 left、right 和最长区间值 l、r
	left, right, maxLeft, maxRight := 0, 0, 0, 0
	for i := 1; i < length; i++ {
		right++

		// 前面连续数字和小于当前数，则重新开始
		if dp[i-1]+array[i] < array[i] {
			dp[i] = array[i]
			left = right
		} else {
			dp[i] = dp[i-1] + array[i]
		}

		if dp[i] > max || dp[i] == max && (right-left) > (maxRight-maxLeft) {
			max = dp[i]
			maxLeft = left
			maxRight = right
		}
	}

	return array[maxLeft : maxRight+1]
}

func main() {
	arr := []int{1, -2, 3, 10, -4, 7, 2, -5}
	ret := findGreatestSumOfSubArr(arr)
	fmt.Println(ret)
}
