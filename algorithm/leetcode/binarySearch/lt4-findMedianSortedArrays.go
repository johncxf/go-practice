// [L4-困难] 寻找两个正序数组的中位数
package main

import (
	"fmt"
)

// 这里没有用二分，并没有达到题目要求的时间复杂度
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 先合并两个有序数组
	nums := mergeSortedArrays(nums1, nums2)
	count := len(nums)
	if count%2 == 0 {
		return float64(nums[count/2-1]+nums[count/2]) / 2.0
	} else {
		return float64(nums[count/2])
	}
}

func mergeSortedArrays(l, r []int) (res []int) {
	for len(l) != 0 && len(r) != 0 {
		if l[0] <= r[0] {
			res = append(res, l[0])
			l = l[1:]
		} else {
			res = append(res, r[0])
			r = r[1:]
		}
	}
	for len(l) != 0 {
		res = append(res, l[0])
		l = l[1:]
	}
	for len(r) != 0 {
		res = append(res, r[0])
		r = r[1:]
	}
	return
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
