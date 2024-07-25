// [L62-中等] 不同路径
package main

import "fmt"

// 动态规划
func uniquePaths(m int, n int) int {
	// 构建dp表
	// dp[i][j] 到达 i,j 位置的路径数
	dp := make([][]int, m)
	// 边界处理
	// dp[0][0] = 1
	// dp[i][0] = 1
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	// dp[0][j] = 1
	for j := 1; j < n; j++ {
		dp[0][j] = 1
	}
	// 状态转移
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(2, 3))
	fmt.Println(uniquePaths(3, 3))
}
