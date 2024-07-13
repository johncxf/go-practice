package main

import "fmt"

// 回溯算法
func backtracePermute(nums []int, state []int, selected []bool, res [][]int) [][]int {
	// 当状态长度等于元素数量时，记录解
	if len(nums) == len(state) {
		res = append(res, append([]int{}, state...))
		return res
	}
	// 遍历所有选择
	for i := 0; i < len(nums); i++ {
		if selected[i] {
			continue
		}
		// 更新状态
		selected[i] = true
		state = append(state, nums[i])
		// 进行下一轮选择
		res = backtracePermute(nums, state, selected, res)
		// 回退：撤销选择，恢复到之前的状态
		selected[i] = false
		state = state[:len(state)-1]
	}
	return res
}

// [L46-简单] 全排列
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, len(nums))
	return backtracePermute(nums, state, selected, res)
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	//fmt.Println(permute([]int{0, 1}))
	//fmt.Println(permute([]int{1}))
}
