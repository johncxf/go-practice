// [L4-困难] 寻找两个正序数组的中位数
package main

import (
	"fmt"
)

// 先合并2两个数组，再取中位数
// 时间复杂度：O(N+M)
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	nums := mergeSortedArrays(nums1, nums2)
	n := len(nums)
	if n%2 != 0 {
		return float64(nums[n/2])
	} else {
		return float64(nums[n/2-1]+nums[n/2]) / 2
	}
}

// 合并两个递增数组
func mergeSortedArrays(nums1, nums2 []int) []int {
	var res []int
	for len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] < nums2[0] {
			res = append(res, nums1[0])
			nums1 = nums1[1:]
		} else {
			res = append(res, nums2[0])
			nums2 = nums2[1:]
		}
	}
	if len(nums1) > 0 {
		res = append(res, nums1...)
	}
	if len(nums2) > 0 {
		res = append(res, nums2...)
	}
	return res
}

// 在归并合并基础上进行优化
// 时间复杂度为：O((N+M)/2)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	// 总数
	count := len(nums1) + len(nums2)
	// 左右两个数
	left, right := 0, 0
	// 遍历前count/2个数
	// 当count为奇数时，中位数就是第count/2个
	// 当count为偶数时，中位数就是第count/2和count/2-1两个数的平均数
	for i := 0; i <= count/2; i++ {
		// 左指针移动一位
		left = right
		// 每次取最小的数，并在数组中移除该数
		if len(nums1) > 0 && len(nums2) > 0 {
			if nums1[0] < nums2[0] {
				right = nums1[0]
				nums1 = nums1[1:]
			} else {
				right = nums2[0]
				nums2 = nums2[1:]
			}
		} else if len(nums1) == 0 {
			right = nums2[0]
			nums2 = nums2[1:]
		} else if len(nums2) == 0 {
			right = nums1[0]
			nums1 = nums1[1:]
		}
	}
	if count%2 != 0 {
		return float64(right)
	} else {
		return float64(left+right) / 2
	}
}

func main() {
	fmt.Println(findMedianSortedArrays1([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays1([]int{1, 2}, []int{3, 4}))

	fmt.Println(findMedianSortedArrays2([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays2([]int{1, 2}, []int{3, 4}))
}
