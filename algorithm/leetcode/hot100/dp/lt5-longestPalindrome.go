// [L5-中等] 最长回文子串
package main

import "fmt"

// 动态规划
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	// 构建dp表：dp[i][j] 表示 s[i..j] 是否是回文串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// 最大长度和开始下标
	max, index := 1, 0
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 { // 长度小于2，1个字母为回文串，2个字母相等也是回文串
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 更新回文长度和起始位置
			if dp[i][j] == true && j-i+1 > max {
				max = j - i + 1
				index = i
			}
		}
	}
	return s[index : index+max]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return left + 1, right - 1
}

// 中心扩散法
func longestPalindrome2(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	left, right := 0, 0
	for i := 0; i < n-1; i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > right-left {
			left, right = left1, right1
		}
		if right2-left2 > right-left {
			left, right = left2, right2
		}
	}

	return s[left : right+1]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))

	fmt.Println(longestPalindrome2("babad"))
	fmt.Println(longestPalindrome2("cbbd"))
}
