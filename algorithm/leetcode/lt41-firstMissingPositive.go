// [L41-困难] 缺失的第一个正数
package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 将数组中所有小于等于 0 的数修改为 N+1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	// 定义取绝对值函数
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	// 遍历所有元素
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		// 将所有小于 n 的元素打标记
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	// 取第一个正数 + 1，如果都是负数，那答案就是 n + 1
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}
