// [L84-困难] 柱状图中最大的矩形
package main

import (
	"fmt"
)

// 暴力-枚举宽
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	maxArea := heights[0]
	// 遍历数组 heights，以每个位置为左边界
	for l := 0; l < n; l++ {
		// 最小高度
		minHeight := heights[l]
		// 右边界
		for r := l; r < n; r++ {
			if heights[r] < minHeight {
				minHeight = heights[r]
			}
			area := (r - l + 1) * minHeight
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// 暴力-枚举高
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	maxArea := heights[0]
	// 以当前元素为矩阵最低高度，向两侧进行枚举
	for i := 0; i < n; i++ {
		h := heights[i]
		l, r := i, i
		// 向左、右两侧侧遍历，遇到小于当前高度h终止
		for l-1 >= 0 && heights[l-1] >= h {
			l--
		}
		for r+1 < n && heights[r+1] >= h {
			r++
		}
		area := (r - l + 1) * h
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// 优化1
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func largestRectangleArea3(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	// 单调栈：单调递增
	stack := make([]int, 0)
	// 从左到右遍历所有元素，找出所有i左侧最近低于heights[i]的元素坐标
	for i := 0; i < n; i++ {
		// 保持单调栈属性：栈顶元素不能大于等于当前元素，不然进行出栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			// 栈为空，说明左侧不存在小于当前高度的元素，置为-1
			left[i] = -1
		} else {
			// 不为空，取最近的小于当前高度的坐标
			left[i] = stack[len(stack)-1]
		}
		// 当前位置入栈
		stack = append(stack, i)
	}
	// 单调栈重置
	stack = make([]int, 0)
	// 从右往左遍历所有元素，找出所有i右侧最近低于heights[i]的元素坐标
	for i := n - 1; i >= 0; i-- {
		// 保持单调栈属性：栈顶元素不能大于等于当前元素，不然进行出栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			// 栈为空，说明右侧没有小于它的元素，置为 n
			right[i] = n
		} else {
			// 不为空，取最近的小于当前高度的坐标
			right[i] = stack[len(stack)-1]
		}
		// 当前位置入栈
		stack = append(stack, i)
	}
	maxArea := heights[0]
	// 遍历所有位置 i，以 heights[i] 为高度计算面积
	for i := 0; i < n; i++ {
		area := (right[i] - left[i] - 1) * heights[i]
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// 优化2
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func largestRectangleArea4(heights []int) int {
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
	fmt.Println(largestRectangleArea3([]int{2, 1, 5, 6, 2, 3}))

	fmt.Println(largestRectangleArea4([]int{2, 1, 5, 6, 2, 3}))
}
