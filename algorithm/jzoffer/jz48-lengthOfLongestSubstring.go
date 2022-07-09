package main

import (
	"fmt"
	"math"
)

/**
 * 滑动窗口
 *
 * @param s string 字符串
 * @return int 整型
 */
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 1 {
		return 1
	}

	hash := make([]int, 256)
	left, right, max := 0, 0, 1
	for right < n {
		hash[s[right]] += 1
		for hash[s[right]] != 1 {
			hash[s[left]]--
			left++
		}
		max = int(math.Max(float64(max), float64(right-left+1)))
		right++
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
