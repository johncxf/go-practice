// [L2414-中等] 最长的字母序连续子字符串的长度
package main

import "fmt"

// 枚举
// - 时间复杂度：O(N^2)
// - 空间复杂度：O(1)
func longestContinuousSubstring1(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		count := 1
		for j := i; j < len(s)-1 && s[j]+1 == s[j+1]; j++ {
			count++
		}
		if count > max {
			max = count
		}
	}
	return max
}

// 模拟
// - 时间复杂度：O(N)
// - 空间复杂度：O(1)
func longestContinuousSubstring2(s string) int {
	max := 0
	// 定义左右区间
	l, r := 0, 1
	// 遍历数组
	for ; r < len(s); r++ {
		// 不连续则更新当前最大值
		if s[r-1]+1 != s[r] {
			if r-l > max {
				max = r - l
			}
			// 左区间更新
			l = r
		}
	}
	if r-l > max {
		max = r - l
	}
	return max
}

func main() {
	fmt.Println(longestContinuousSubstring1("abacaba"))
	fmt.Println(longestContinuousSubstring1("abcde"))

	fmt.Println(longestContinuousSubstring2("abacaba"))
	fmt.Println(longestContinuousSubstring2("abcde"))
}
