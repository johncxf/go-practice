// [L7363-中等] 划分字母区间
package main

import "fmt"

func partitionLabels(s string) []int {
	// 遍历字符串，找出每个字母最后出现的位置
	lastIndex := [26]int{}
	for i, v := range s {
		lastIndex[v-'a'] = i
	}
	res := make([]int, 0)
	start, end := 0, 0
	// 遍历字符串
	for i, v := range s {
		// 更新区间最靠后的位置
		if lastIndex[v-'a'] > end {
			end = lastIndex[v-'a']
		}
		// 达到区间末尾，更新结果
		if i == end {
			res = append(res, end-start+1)
			// 更新下一段区间的开始下标
			start = end + 1
		}
	}
	return res
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}
