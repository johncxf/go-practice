// [L35-简单] 搜索插入位置
package main

import "fmt"

// 二分查找
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if target == nums[mid] { // 找到
			return mid
		} else if target < nums[mid] { // 目标值小于中间值，且找不到
			right = mid - 1
		} else { // 目标值大于中间值，且找不到
			left = mid + 1
		}
	}
	return left
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
