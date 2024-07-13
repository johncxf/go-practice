// [L11-中等] 盛最多水的容器
package main

import (
	"fmt"
)

// 双指针
func maxArea(height []int) int {
	n := len(height)
	left, right, max := 0, n-1, 0
	for left < right {
		tmp := 0
		if height[left] < height[right] {
			tmp = height[left] * (right - left)
			left++
		} else {
			tmp = height[right] * (right - left)
			right--
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
