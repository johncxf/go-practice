package main

import "fmt"

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	// 获取元素和，和最大元素
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	// 若元素和为奇数，则不可能分为两个相等的子集
	if sum%2 != 0 {
		return false
	}
	// 若最大元素大于总和的一半，也不可能分为两个相等的子集
	target := sum / 2
	if max > target {
		return false
	}

	// 构建 dp 表
	dp := make([][]bool, n)
	for i := range nums {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[n-1][target]
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
