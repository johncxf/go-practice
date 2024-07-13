// [L131-中等] 分割回文串
package main

import "fmt"

func partition(s string) [][]string {
	path := make([]string, 0)
	ans := make([][]string, 0)
	backtrackPa(s, path, 0, &ans)
	return ans
}

func backtrackPa(s string, path []string, start int, ans *[][]string) {
	// 所有元素都遍历完成，将path添加到结果中
	if start == len(s) {
		*ans = append(*ans, append([]string{}, path...))
		return
	}
	// 遍历所有选择
	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
			backtrackPa(s, path, i+1, ans)
			path = path[:len(path)-1]
		}
	}
	return
}

// 判断是否是回文串
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}
