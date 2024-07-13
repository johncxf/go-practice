// [L139-中等] 单词拆分
package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	// 将字典转化为 hashMap
	hash := make(map[string]bool)
	for _, v := range wordDict {
		hash[v] = true
	}
	// 构建 dp 表
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && hash[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
