// [L39-中等] 组合总和
package main

import (
	"fmt"
	"sort"
)

// 回溯
func backtraceCombinationSum(start, target int, state, choices []int, res *[][]int) {
	// 子集和等于 target 时，记录解
	if target == 0 {
		*res = append(*res, append([]int{}, state...))
		return
	}
	// 遍历所有选择
	// 剪枝二：从 start 开始遍历，避免生成重复子集
	for i := start; i < len(choices); i++ {
		// 剪枝一：若子集和超过 target ，则直接结束循环
		// 这是因为数组已排序，后边元素更大，子集和一定超过 target
		if target-choices[i] < 0 {
			break
		}
		// 更新状态
		state = append(state, choices[i])
		// // 进行下一轮选择
		backtraceCombinationSum(i, target-choices[i], state, choices, res)
		// 回退
		state = state[:len(state)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	// 先进行排序，为去重
	sort.Ints(candidates)
	res := make([][]int, 0)
	state := make([]int, 0)
	backtraceCombinationSum(0, target, state, candidates, &res)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
