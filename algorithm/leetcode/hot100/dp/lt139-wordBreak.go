// [L139-中等] 单词拆分
package main

import (
	"fmt"
)

func wordBreak1(s string, wordDict []string) bool {
	n := len(s)
	// 构建dp表
	//dp[i] 表示以 i 结尾的字符串是否可以被 wordDict 中组合而成
	dp := make([]bool, n+1)
	dp[0] = true
	// 状态转移：遍历所有 i
	for i := 1; i <= n; i++ {
		// 进行选择
		for _, word := range wordDict {
			length := len(word)
			// 以 i 结尾的字符串长度需要大于当前单词长度
			// 单词相等
			if i-length >= 0 && s[i-length:i] == word {
				dp[i] = dp[i] || dp[i-length]
			}
		}
	}
	return dp[n]
}

func wordBreak2(s string, wordDict []string) bool {
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
	fmt.Println(wordBreak1("leetcode", []string{"leet", "code"}))

	fmt.Println(wordBreak2("leetcode", []string{"leet", "code"}))
}
