package main

import (
	"fmt"
	"math"
)

/**
 * 动态规划
 *
 * @param grid int整型二维数组
 * @return int整型
 */
func maxValue(grid [][]int) int {
	// n 行，m 列
	n, m := len(grid), len(grid[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 && j > 0 {
				grid[i][j] += int(math.Max(float64(grid[i-1][j]), float64(grid[i][j-1])))
			} else if i > 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			} else if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
			}
		}
	}

	return grid[n-1][m-1]
}

func main() {
	matrix := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(maxValue(matrix))
}
