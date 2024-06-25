package main

import "fmt"

// [L55-中等] 跳跃游戏
func canJump(nums []int) bool {
	k, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if i > k {
			return false
		}
		if i+nums[i] > k {
			k = i + nums[i]
		}
		if k >= n-1 {
			return true
		}
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
