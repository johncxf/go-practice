package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	// 排序
	sort.Ints(candidates)
	// 过程数组
	state := make([]int, 0)
	// 结果数组
	res := make([][]int, 0)
	return backtrace(candidates, target, state, 0, res)
}

func backtrace(candidates []int, target int, state []int, start int, res [][]int) [][]int {
	// 为0时符合，添加子集到结果集中
	if target == 0 {
		//res = append(res, state)
		res = append(res, append([]int{}, state...))
		return res
	}
	// 遍历，从 start 开始，避免重复子集
	for i := start; i < len(candidates); i++ {
		// 元素和不能超过 target
		if target-candidates[i] < 0 {
			break
		}
		// 更新子集
		state = append(state, candidates[i])
		// 递归，进行下一轮
		res = backtrace(candidates, target-candidates[i], state, i, res)
		// 回退
		state = state[:len(state)-1]
	}
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
