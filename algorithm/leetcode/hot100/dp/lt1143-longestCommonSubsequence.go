// [L1143-中等] 最长公共子序列
package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// 构建dp表
	// dp[i][j] 表示text1前i个字符，text2前j个字符的最长公共子序列
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 边界确定
	// 当i或j为0是对应值为0
	// dp[i][0] = 0
	// dp[0][j] = 0
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}
	// 每增加一个字符，text1[i-1]和text2[j-1]是否相同
	// text1[i-1] == text2[j-1]: dp[i][j] = dp[i-1][j-1] + 1
	// text1[i-1] != text2[j-1]: dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
