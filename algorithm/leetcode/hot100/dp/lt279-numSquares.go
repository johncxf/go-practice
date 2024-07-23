// [L279-中等] 完全平方数
package main

import (
	"fmt"
)

// 动态规划
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 设置一个较大数
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			if dp[i-j*j]+1 < dp[i] {
				dp[i] = dp[i-j*j] + 1
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
}
