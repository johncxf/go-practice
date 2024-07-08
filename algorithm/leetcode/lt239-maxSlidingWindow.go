// [L239-困难] 滑动窗口最大值
package main

import "fmt"

// 超时 - 不适用
func maxSlidingWindow99(nums []int, k int) []int {
	maxIndex, max := 0, nums[0]
	ans, kArr := make([]int, 0), make([]int, 0)
	for i := 0; i < k; i++ {
		kArr = append(kArr, nums[i])
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}
	ans = append(ans, max)
	for i := k; i < len(nums); i++ {
		kArr = append(kArr, nums[i])
		kArr = kArr[1:]
		if nums[i] > max {
			max = nums[i]
			maxIndex = len(kArr) - 1
		} else {
			if maxIndex > 0 {
				maxIndex--
			} else {
				max = kArr[0]
				maxIndex = 0
				for j, v := range kArr {
					if v > max {
						maxIndex = j
						max = v
					}
				}
			}
		}
		ans = append(ans, max)
	}
	return ans
}

func maxSlidingWindow(nums []int, k int) []int {
	q := make([]int, 0)
	push := func(i int) {
		// 删除队列中所有大于 num[i] 的元素
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	// 先将前 k 个元素入队
	for i := 0; i < k; i++ {
		push(i)
	}

	ans := make([]int, 0)
	ans = append(ans, nums[q[0]])
	for i := k; i < len(nums); i++ {
		// 入队
		push(i)
		// 保证队列长度小于 k
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	//fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}
