// [L72-中等] 编辑距离
package main

import "fmt"

// 动态规划1
func minDistance1(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	// 构建 dp 表
	// dp[i][j]：word1的前i个字符变为word2的前j个字符所需要的最小距离
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	// 边界确认
	// dp[0][0] = 0，默认就是0，可以省略
	// dp[i][0] = i
	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}
	// dp[0][j] = j
	for j := 1; j <= m; j++ {
		dp[0][j] = j
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 尾部字符相等
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 取最小值
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > dp[i][j-1]+1 {
					dp[i][j] = dp[i][j-1] + 1
				}
				if dp[i][j] > dp[i-1][j]+1 {
					dp[i][j] = dp[i-1][j] + 1
				}
			}
		}
	}
	return dp[n][m]
}

// 动态规划2
func minDistance2(word1 string, word2 string) int {
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
	fmt.Println(minDistance1("horse", "ros"))

	fmt.Println(minDistance2("horse", "ros"))
}
