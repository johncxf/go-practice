package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	return backtrace1(nums, state, res)
}

func backtrace1(nums []int, state []int, res [][]int) [][]int {
	if len(nums) == len(state) {
		res = append(res, append([]int{}, state...))
		return res
	}
	for i := 0; i < len(nums); i++ {
		if inArray(state, nums[i]) {
			continue
		}
		state = append(state, nums[i])
		res = backtrace1(nums, state, res)
		state = state[:len(state)-1]
	}
	return res
}

func inArray(arr []int, target int) bool {
	exist := false
	for _, v := range arr {
		if target == v {
			exist = true
			break
		}
	}
	return exist
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}
