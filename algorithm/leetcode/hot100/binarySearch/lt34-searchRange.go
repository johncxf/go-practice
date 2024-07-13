package main

import "fmt"

// [L34-中等] 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			tmp := mid
			for mid >= left && target == nums[mid] {
				mid--
			}
			res[0] = mid + 1
			mid = tmp
			for mid <= right && target == nums[mid] {
				mid++
			}
			res[1] = mid - 1
			break
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{1}, 1))

}
