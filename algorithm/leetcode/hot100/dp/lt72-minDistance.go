// [L72-中等] 编辑距离
package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// 有一个字符串为空
	if m == 0 || n == 0 {
		return m + n
	}
	// 构建 dp 表
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	// 状态转移
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			left_down := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				left_down++
			}
			// 取最小
			dp[i][j] = left
			if down < dp[i][j] {
				dp[i][j] = down
			}
			if left_down < dp[i][j] {
				dp[i][j] = left_down
			}
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
}
