// [L5-中等] 最长回文子串
package main

import "fmt"

// 动态规划
func longestPalindrome(s string) string {
	n := len(s)
	// 构建dp表
	// dp[l][r] 表示字符串 l 到 r 区间字符是否为回文串，即 s[l:r+1] 是否为回文串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// 边界确定
	dp[0][0] = true
	// 状态转移
	// 当 s[l] == s[r]，若 r-l <= 3 则 dp[l][r] = true，若 r-l > 3 则 dp[l][r] = dp[l+1][r-1]
	// 当 s[l] != s[r] 则: dp[i][j] = false
	// 最大长度，最大长度开始下标
	max, index := 1, 0
	for r := 1; r < n; r++ {
		for l := 0; l < r; l++ {
			if s[l] != s[r] {
				dp[l][r] = false
			} else {
				if r-l < 3 {
					// s[l] 和 s[r] 相等时，字符串为1、2、3个字符时都是回文串
					dp[l][r] = true
				} else {
					dp[l][r] = dp[l+1][r-1]
				}
			}
			if dp[l][r] && r-l+1 > max {
				max = r - l + 1
				index = l
			}
		}
	}
	return s[index : index+max]
}

// 中心扩散法
func longestPalindrome2(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	left, right := 0, 0
	for i := 0; i < n-1; i++ {
		// 以 i 为中心进行左右扩散
		left1, right1 := expandAroundCenter(s, i, i)
		if right1-left1 > right-left {
			left, right = left1, right1
		}
		// 以 i, i+1 为中心进行左右扩散（2个连续字符相等的场景）
		if s[i] == s[i+1] {
			left2, right2 := expandAroundCenter(s, i, i+1)
			if right2-left2 > right-left {
				left, right = left2, right2
			}
		}
	}
	return s[left : right+1]
}

// 进行左右扩散
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

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))

	fmt.Println(longestPalindrome2("babad"))
	fmt.Println(longestPalindrome2("cbbd"))
}
