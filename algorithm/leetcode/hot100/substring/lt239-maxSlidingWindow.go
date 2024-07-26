// [L239-困难] 滑动窗口最大值
package main

import "fmt"

// 暴力，超时
func maxSlidingWindow1(nums []int, k int) []int {
	n := len(nums)
	var ans []int
	// 以i为窗口起始点进行遍历
	for i := 0; i <= n-k; i++ {
		max := nums[i]
		// 查找 i+k 个元素中的最大值
		for j := i + 1; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		ans = append(ans, max)
	}
	return ans
}

func maxSlidingWindow2(nums []int, k int) []int {
	// 单调队列，单调递减
	q := make([]int, 0)
	// 实现单调队列入队操作
	push := func(i int) {
		// 删除所有小于 num 的元素，并将 num 添加至队尾，保持单调递减
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	// 先将前 k 个元素入队
	for i := 0; i < k; i++ {
		push(i)
	}
	// 结果集合
	ans := make([]int, 0)
	ans = append(ans, nums[q[0]])
	// 移动窗口右坐标 i
	for i := k; i < len(nums); i++ {
		// 当前元素入队
		push(i)
		// 如果当前队头元素不在长度为 k 的区间，则删除该队头元素
		for q[0] <= i-k {
			q = q[1:]
		}
		// 将队头加入结果集合
		ans = append(ans, nums[q[0]])
	}
	return ans
}

func main() {
	fmt.Println(maxSlidingWindow1([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

	fmt.Println(maxSlidingWindow2([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow2([]int{1, -1}, 1))
}
