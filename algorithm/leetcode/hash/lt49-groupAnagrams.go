package main

import (
	"fmt"
	"sort"
)

// [L49-中等] 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	hash := map[string][]string{}
	for _, str := range strs {
		// 将字符串转化成数组
		s := []byte(str)
		// 对字符串数组进行排序
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		// 将字符串数组中的元素类型进行转化
		sortedStr := string(s)
		// 以字符串为key构建map
		hash[sortedStr] = append(hash[sortedStr], str)
	}
	// 重新组装成数组
	ans := make([][]string, 0, len(hash))
	for _, v := range hash {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
