// [L53-中等] 最大子数组和
package main

import "fmt"

// 动态规划
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// 迭代
func maxSubArray2(nums []int) int {
	// 最大值，当前子数组和
	max, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > cur+nums[i] {
			// 当前元素值大于子数组和，则更新当前子数组和值
			cur = nums[i]
		} else {
			// 当前子数组和值需要加上当前元素
			cur += nums[i]
		}
		if cur > max {
			max = cur
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))

	fmt.Println(maxSubArray2([]int{5, 4, -1, 7, 8}))
}
