// [L84-困难] 柱状图中最大的矩形
package main

import (
	"fmt"
	"math"
)

// 暴力-枚举宽
// 时间复杂度：O(N^2)，超时
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	ans := 0
	// 左侧边界
	for i := 0; i < n; i++ {
		min := math.MaxInt32
		// 右侧边界
		for j := i; j < n; j++ {
			if heights[j] < min {
				min = heights[j]
			}
			// 计算面积
			tmp := (j - i + 1) * min
			// 更新面积
			if tmp > ans {
				ans = tmp
			}
		}
	}
	return ans
}

// 暴力-枚举高
// 时间复杂度：O(N^2)，超时
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	ans := 0
	// 遍历数组，以每个元素作为矩阵的高h，向两侧延伸，直到遇到高度小于h的柱子后停止
	for mid := 0; mid < n; mid++ {
		h := heights[mid]
		left, right := mid, mid
		// 确定左边界，遇到高度小于h的柱子后停止
		for left-1 >= 0 && heights[left-1] >= h {
			left--
		}
		// 确定右边界
		for right+1 < n && heights[right+1] >= h {
			right++
		}
		// 计算面积
		area := (right - left + 1) * h
		if area > ans {
			ans = area
		}
	}
	return ans
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	stack := []int{}
	// 遍历数组
	for i := 0; i < n; i++ {
		// 栈不为空，且当前高度小于等于前一个高度时
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			// 将当前坐标加入右侧坐标数组
			right[stack[len(stack)-1]] = i
			// 出栈
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i]--
		} else {
			// 将当前坐标加入左侧坐标数组
			left[i] = stack[len(stack)-1]
		}
		// 入栈
		stack = append(stack, i)
	}
	ans := 0
	// 根据左右边界值计算面积，获得最大面积
	for i := 0; i < n; i++ {
		area := (right[i] - left[i] - 1) * heights[i]
		if area > ans {
			ans = area
		}
	}
	return ans
}

func main() {
	fmt.Println(largestRectangleArea1([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea2([]int{2, 1, 5, 6, 2, 3}))

	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
