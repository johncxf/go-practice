// [L49-中等] 字母异位词分组
package main

import (
	"fmt"
	"sort"
)

// 排序+哈希表
// - 时间复杂度：O(nklogk)
// - 空间复杂度：O(nk)
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

// 计数+哈希表
// - 时间复杂度：O(n(k+∣Σ∣))，n 是 strs 中的字符串的数量，k 是 strs 中的字符串的的最大长度，Σ 是字符集
// - 空间复杂度：O(n(k+∣Σ∣))
func groupAnagrams2(strs []string) [][]string {
	hash := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, ch := range str {
			cnt[ch-'a']++
		}
		hash[cnt] = append(hash[cnt], str)
	}
	var ans [][]string
	for _, v := range hash {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	fmt.Println(groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
