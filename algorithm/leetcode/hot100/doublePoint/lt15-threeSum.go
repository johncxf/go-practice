// [L15-中等] 三数之和
package main

import (
	"fmt"
	"sort"
)

// 排序+双指针
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	ret := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		// 跳过重复数字
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		need := 0 - nums[i]
		for left < right {
			if need == nums[left]+nums[right] {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				// 跳过相同的数字
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if need > nums[left]+nums[right] {
				left++
			} else {
				right--
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
