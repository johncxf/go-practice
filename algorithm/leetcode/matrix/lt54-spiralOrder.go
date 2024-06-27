// [L54-中等] 螺旋矩阵
package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	// 定义4个边界
	upper, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	ans := make([]int, 0)
	for {
		// 向右移动直到最右
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[upper][i])
		}
		// 重新设定上边界
		upper++
		if upper > down {
			break
		}
		// 向下
		for i := upper; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		// 重新设定右边界
		right--
		if right < left {
			break
		}
		// 向左
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		// 重新设定下边界
		down--
		if down < upper {
			break
		}
		// 向上
		for i := down; i >= upper; i-- {
			ans = append(ans, matrix[i][left])
		}
		// 重新设定左边界
		left++
		if left > right {
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
