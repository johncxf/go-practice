// [L41-困难] 缺失的第一个正数
package main

import (
	"fmt"
)

// 哈希表
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func firstMissingPositive1(nums []int) int {
	hash := make(map[int]bool, 0)
	// 遍历数组，将数组中所有数存入 hash 表
	for _, num := range nums {
		hash[num] = true
	}
	n := len(nums)
	// 从1开始枚举正整数，遇到hash表中不存在的就是最小的正整数
	for i := 1; i <= n; i++ {
		if !hash[i] {
			return i
		}
	}
	// 如果 1～n 都存在数组中，那最小的正整数就是 n+1
	return n + 1
}

// 原地置换
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func firstMissingPositive2(nums []int) int {
	n := len(nums)
	// 遍历一遍数组，将 nums[i] 放到下标为 nums[i]-1 的位置上
	for i := 0; i < n; i++ {
		// nums[i] 范围在 1 ～ n
		// 这里用 for 循环是因为当前位置 i 与 nums[i]-1 位置换了元素，那当前位置 i 需要重新进行判断
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
		// 这样也可以
		//if nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
		//	nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		//	i--
		//}
	}
	// 再次遍历数组，找出 nums[i] 不等于 i-1 的数
	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	fmt.Println(firstMissingPositive1([]int{1, 2, 0}))

	fmt.Println(firstMissingPositive2([]int{1, 2, 0}))
}
