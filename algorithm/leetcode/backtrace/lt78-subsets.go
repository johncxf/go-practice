package main

import "fmt"

// 回溯
func backtrackSubsets(nums []int, path []int, start int, res *[][]int) {
	// 记录结果
	*res = append(*res, append([]int{}, path...))
	// 遍历所有选择
	for i := start; i < len(nums); i++ {
		// 更新状态
		path = append(path, nums[i])
		// 进行下一次选择
		backtrackSubsets(nums, path, i+1, res)
		// 回退
		path = path[:len(path)-1]
	}
}

// [L78-中等] 子集
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	backtrackSubsets(nums, path, 0, &res)
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
