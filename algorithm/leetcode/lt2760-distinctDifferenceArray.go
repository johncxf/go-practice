package main

import "fmt"

func distinctDifferenceArray(nums []int) {
	var ret []int
	n := len(nums)
	for i := 0; i < n; i++ {
		//fmt.Println(nums[i])
		ret = append(ret, nums[i])
	}
	fmt.Println(ret)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	distinctDifferenceArray(arr)
}
