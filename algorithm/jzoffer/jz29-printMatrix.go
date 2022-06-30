package main

import "fmt"

/**
 * 边界模拟法
 *
 * @param matrix int 整型二维数组
 * @return int 整型一维数组
 */
func printMatrix(matrix [][]int) []int {
	result := []int{}
	n := len(matrix)
	if 0 == n {
		return result
	}

	// 左、右、上、下 边界
	left, right, up, down := 0, len(matrix[0])-1, 0, n-1
	for left <= right && up <= down {
		// 遍历上边界从左到右
		for i := left; i <= right; i++ {
			result = append(result, matrix[up][i])
		}
		// 上边界向下
		up++
		if up > down {
			break
		}

		// 右边界的从上到下
		for i := up; i <= down; i++ {
			result = append(result, matrix[i][right])
		}
		//右边界向左
		right--
		if left > right {
			break
		}

		// 下边界的从右到左
		for i := right; i >= left; i-- {
			result = append(result, matrix[down][i])
		}
		// 下边界向上
		down--
		if up > down {
			break
		}

		// 左边界从下向上
		for i := down; i >= up; i-- {
			result = append(result, matrix[i][left])
		}
		// 左边界向右
		left++
		if left > right {
			break
		}
	}

	return result
}

func main() {
	ret := printMatrix([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})
	fmt.Println(ret)
}
