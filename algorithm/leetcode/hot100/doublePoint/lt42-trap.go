// [L42-困难] 接雨水
package main

import "fmt"

// 暴力枚举
// 时间复杂度：O(N^2)
func trap1(height []int) int {
	n := len(height)
	ans := 0
	for i := 1; i < n-1; i++ {
		// 向左查找大于本身高度的最大值
		leftMax := 0
		for l := i - 1; l >= 0; l-- {
			if height[l] > height[i] && height[l] > leftMax {
				leftMax = height[l]
			}
		}
		// 向右查找大于本身高度的最大值
		rightMax := 0
		for r := i + 1; r < n; r++ {
			if height[r] > height[i] && height[r] > rightMax {
				rightMax = height[r]
			}
		}
		// 左右都有大于当前高度的柱子才能接到雨水
		if leftMax != 0 && rightMax != 0 {
			// 接的雨水等于两个最高柱子的最小值减去自身柱子
			min := leftMax
			if rightMax < min {
				min = rightMax
			}
			ans += min - height[i]
		}
	}
	return ans
}

// 动态规划
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func trap2(height []int) int {
	n := len(height)
	// 建立dp表
	leftMax, rightMax := make([]int, n), make([]int, n)
	// 边界确认
	leftMax[0], rightMax[n-1] = height[0], height[n-1]
	// 状态转移
	for i := 1; i < n; i++ {
		leftMax[i] = height[i]
		if leftMax[i-1] > leftMax[i] {
			leftMax[i] = leftMax[i-1]
		}
	}
	for j := n - 2; j >= 0; j-- {
		rightMax[j] = height[j]
		if rightMax[j+1] > rightMax[j] {
			rightMax[j] = rightMax[j+1]
		}
	}
	// 遍历柱子，计算雨水总量
	ans := 0
	for i := 1; i < n-1; i++ {
		// 取较小值
		if leftMax[i] < rightMax[i] {
			ans += leftMax[i] - height[i]
		} else {
			ans += rightMax[i] - height[i]
		}
	}
	return ans
}

// 双指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func trap3(height []int) int {
	n := len(height)
	if n == 0 {
		return -1
	}
	ans := 0
	left, right, leftMax, rightMax := 0, n-1, 0, 0
	for left < right {
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}
		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

func main() {
	fmt.Println(trap1([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap3([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
