// [L240-中等] 搜索二维矩阵II
package main

import "fmt"

// 暴力
func searchMatrix1(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, col := range row {
			if col == target {
				return true
			}
		}
	}
	return false
}

// Z 字查找
func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix1([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))

	fmt.Println(searchMatrix2([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
}
