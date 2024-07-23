// [L198-中等] 打家劫舍
package main

import "fmt"

// 动态规划：dp[i]=max(dp[i-1],dp[i-2]+nums[i])，i >= 2
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}
	for i := 2; i < n; i++ {
		if dp[i-2]+nums[i] > dp[i-1] {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
