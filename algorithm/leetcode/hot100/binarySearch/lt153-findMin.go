// [L153-中等] 搜索旋转排序数组中的最小值
package main

import "fmt"

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	min := nums[0]
	for l <= r {
		mid := (l + r) / 2
		// 中间值小于等于最右侧值，这部分序列左侧mid即是这部分的最小值，记录下
		if nums[mid] <= nums[r] {
			if nums[mid] < min {
				min = nums[mid]
			}
			r = mid - 1
		} else {
			// 中间值大于最右侧值，说明这部分序列中存在翻转点，即最小值
			l = mid + 1
		}
	}
	return min
}

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}
