// [L55-中等] 跳跃游戏
package main

import "fmt"

func canJump(nums []int) bool {
	k, n := 0, len(nums)
	for i := 0; i < n; i++ {
		// i 之前可达到最远位置没能到达当前i位置，说明无法达到
		if i > k {
			return false
		}
		// 取当前可达到的最远位置
		if i+nums[i] > k {
			k = i + nums[i]
		}
		// 如果当前可达到最远位置已大于 n-1，直接返回 true
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
