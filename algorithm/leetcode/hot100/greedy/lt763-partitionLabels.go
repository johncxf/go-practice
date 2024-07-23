// [L7363-中等] 划分字母区间
package main

import "fmt"

func partitionLabels1(s string) []int {
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

func partitionLabels2(s string) []int {
	// 存在每个字母个数
	hash := [126]int{}
	for _, ch := range s {
		hash[ch]++
	}
	// 结果集
	ans := make([]int, 0)
	// 存储当前区间内存在，但是未全部包含完的字母
	tmp := map[byte]bool{}
	// 当前区间长度
	length := 0
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		ch := s[i]
		// 区间长度+1
		length++
		// 该字母剩余数量-1
		hash[ch]--
		// 将字母加入 tmp 中
		tmp[ch] = true
		if hash[ch] == 0 {
			delete(tmp, ch)
		}
		if len(tmp) == 0 {
			ans = append(ans, length)
			length = 0
		}
	}
	return ans
}

func main() {
	fmt.Println(partitionLabels1("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels1("eccbbbbdec"))

	fmt.Println(partitionLabels2("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels2("eccbbbbdec"))
}
