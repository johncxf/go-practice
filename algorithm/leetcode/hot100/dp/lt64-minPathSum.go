// [L64-中等] 最小路径和
package main

import "fmt"

// 动态规划
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 构建dp表
	// dp[i][j]：达到 i,j 的最小路径
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 边界确定
	// dp[0][0]=grid[0][0]
	dp[0][0] = grid[0][0]
	// 当i为0时，则 dp[i][0] = dp[i-1][0] + grid[i][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// 当j为0时，则 dp[0][j] = dp[0][j-1] + grid[0][j]
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	// 状态转移
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))

}
