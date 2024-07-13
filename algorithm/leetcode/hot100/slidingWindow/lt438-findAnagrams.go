package main

import (
	"fmt"
)

// [L438-中等] 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	n, m := len(s), len(p)
	if n < m {
		return nil
	}

	// 维护2个m个元素的数组，统计每个字符出现的次数
	// 利用golang数组可以使用==比较的特性，判断2个窗口数组是否相等
	var sCount, pCount [123]int
	for i, ch := range p {
		sCount[s[i]]++
		pCount[ch]++
	}

	var ans []int
	if sCount == pCount {
		ans = append(ans, 0)
	}
	for i, ch := range s[:n-m] {
		sCount[ch]--
		sCount[s[i+m]]++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
