package main

import "fmt"

// [L128-中等] 最长连续序列
func longestConsecutive(nums []int) int {
	// 构建 hash 表
	hash := map[int]bool{}
	for _, v := range nums {
		hash[v] = true
	}
	max := 0
	// 遍历 hash 表
	for num := range hash {
		// 只取 num 为最小值的情况
		if !hash[num-1] {
			// 当前值
			current := num
			// 以当前值为最小值情况下的递增序列长度
			count := 1
			// 查找依次增加的序列
			for hash[current+1] {
				current++
				count++
			}
			// 更新最大序列长度值
			if count > max {
				max = count
			}
		}
	}
	return max
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
