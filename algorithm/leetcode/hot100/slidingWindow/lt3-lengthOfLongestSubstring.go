package main

import "fmt"

// [L3-中等] 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	right, max := 0, 0
	// hash map，存储不重复的字符串，用于判断字符串是否出现过
	hash := make(map[byte]int)
	for i := 0; i < n; i++ {
		// 除了 i = 0 第一次，每一次移动左指针的时候都需要删除第一个hash元素
		if i != 0 {
			delete(hash, s[i-1])
		}
		// 不断移动右指针，直到出现重复字符退出
		for right < n && hash[s[right]] == 0 {
			hash[s[right]]++
			right++
		}
		// 更新最大值
		if right-i > max {
			max = right - i
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
