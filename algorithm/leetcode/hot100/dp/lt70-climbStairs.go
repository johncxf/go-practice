// [L70-简单] 爬楼梯
package main

import "fmt"

// 动态规划
func climbStairs(n int) int {
	dp := make([]int, n)
	if n <= 2 {
		return n
	}
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

// 优化空间
func climbStairs2(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	a, b := 1, 2
	// 状态转移：从较小子问题逐步求解较大子问题
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))

	fmt.Println(climbStairs2(1))
	fmt.Println(climbStairs2(2))
	fmt.Println(climbStairs2(3))
}
