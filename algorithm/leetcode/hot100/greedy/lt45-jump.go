// [L45-中等] 跳跃游戏 II
package main

import "fmt"

// 贪心
func jump(nums []int) int {
	// 当前可跳最大距离、步数、当前下标可跳最远距离结束下标
	max, steps, end := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > max {
			max = i + nums[i]
		}
		// 已到结束位置，则更新下一个跳跃
		if i == end {
			end = max
			steps++
			// 当前可跳最远位置大于数组长度则直接推出循环
			if end >= len(nums)-1 {
				break
			}
		}
	}
	return steps
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
	fmt.Println(jump([]int{2, 3, 1, 2, 4, 2, 3}))
}
