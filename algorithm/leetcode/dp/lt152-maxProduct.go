package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	n := len(nums)
	max, min := make([]int, n), make([]int, n)
	max[0], min[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		min[i], max[i] = maxAndMin(max[i-1]*nums[i], min[i-1]*nums[i], nums[i])
	}
	ans := max[0]
	for _, v := range max {
		if v > ans {
			ans = v
		}
	}
	if ans > math.MaxInt32 {
		return 1000000000
	}
	return ans
}

func maxAndMin(a, b, c int) (min, max int) {
	min, max = a, a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{0, 10, 10, 10, 10, 10, 10, 10, 10, 10, -10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 0}))
}
