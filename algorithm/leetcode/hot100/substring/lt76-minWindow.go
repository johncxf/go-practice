// [L76-困难] 最小覆盖子串
package main

import (
	"fmt"
	"math"
)

// 滑动窗口
func minWindow(s string, t string) string {
	n := len(s)
	// hashMap 维护 t 字符以及字符个数
	hashT := make(map[byte]int, 0)
	for i := range t {
		hashT[t[i]]++
	}

	// hashMap 维护窗口中所有字符以及个数
	cnt := make(map[byte]int, 0)

	// 检查窗口是否覆盖 t
	check := func() bool {
		for k, v := range hashT {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	// 子串左右边界坐标
	ansL, ansR := -1, -1
	// 子串最小值
	minLength := math.MaxInt32
	for l, r := 0, 0; r < n; r++ {
		if r < n && hashT[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			// 更新窗口大小
			if r-l+1 < minLength {
				minLength = r - l + 1
				ansL, ansR = l, l+minLength
			}
			// 左指针移动，收缩窗口
			if _, ok := hashT[s[l]]; ok {
				cnt[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))

	hash := map[int]int{
		1: 1,
	}
	fmt.Println(hash[2])
}
