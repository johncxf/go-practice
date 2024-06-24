package main

import "fmt"

// [L238-中等] 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	n := len(nums)
	// 初始化两个数组，分别存储 i 左右两侧的乘机
	lArr, rArr := make([]int, n), make([]int, n)
	lArr[0] = 1
	for i := 1; i < n; i++ {
		lArr[i] = lArr[i-1] * nums[i-1]
	}
	rArr[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rArr[i] = rArr[i+1] * nums[i+1]
	}
	// i 的值为左右两侧乘机
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = lArr[i] * rArr[i]
	}
	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
