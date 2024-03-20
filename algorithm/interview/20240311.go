package main

import "fmt"

// 给定一个字符串s，找出其中重复出现的非空子串的最大长度，即子串的长度大于1。
// 示例输入：s = "abcabcbb"
// 示例输出：3
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	max := 0
	// 构建二维数组
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max
}

// 反转字符串
func reverseString1(str string) string {
	// 第一种
	arr := make([]string, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		arr = append(arr, string(str[i]))
	}
	res := ""
	for _, v := range arr {
		res = res + v
	}
	return res
}

// 反转字符串
func reverseString2(str string) string {
	arr := []rune(str)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring("abcd"))

	//s2 := "Hello World"
	//fmt.Println(reverseString1(s2))
	//fmt.Println(reverseString2(s2))
}
