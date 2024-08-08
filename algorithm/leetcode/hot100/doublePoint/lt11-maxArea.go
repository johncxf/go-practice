// [L11-中等] 盛最多水的容器
package main

import (
	"fmt"
)

// 双指针
func maxArea(height []int) int {
	max, l, r := 0, 0, len(height)-1
	for l < r {
		h := height[l]
		if height[l] > height[r] {
			h = height[r]
			r--
		} else {
			l++
		}
		area := (r - l + 1) * h
		if area > max {
			max = area
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
