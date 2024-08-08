package main

import "fmt"

// [L3-中等] 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	max := 0
	// 记录当前窗口所有元素出现次数
	count := map[byte]int{}
	for l, r := 0, 0; r < len(s); r++ {
		count[s[r]]++
		// 当出现重复元素时，移动左指针，直到没有元素重复
		for count[s[r]] > 1 {
			count[s[l]]--
			l++
		}
		// 计算当前窗口最大值并更新
		if r-l+1 > max {
			max = r - l + 1
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
