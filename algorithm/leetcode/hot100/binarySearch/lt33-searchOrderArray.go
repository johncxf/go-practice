// [L33-中等] 搜索旋转排序数组
package main

import "fmt"

// 二分查找
func searchOrderArray(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] >= nums[left] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(searchOrderArray([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(searchOrderArray([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(searchOrderArray([]int{1}, 0))
	fmt.Println(searchOrderArray([]int{5, 1, 3}, 3))

}
