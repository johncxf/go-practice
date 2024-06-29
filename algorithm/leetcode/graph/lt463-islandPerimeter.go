// [L463-中等] 岛屿的周长
package main

import "fmt"

func dfsPerimeter(grid [][]int, r, c int) int {
	// 超出边界，刚好是边长，返回1
	if r < 0 || len(grid) <= r || c < 0 || len(grid[0]) <= c {
		return 1
	}
	// 如果这个格子是海洋格子，也是边长，返回1
	if grid[r][c] == 0 {
		return 1
	}
	// 如果这个格子是已经遍历过的岛屿，返回0
	if grid[r][c] == 2 {
		return 0
	}
	// 将遍历过的格子标记为2
	grid[r][c] = 2
	// 访问上、下、左、右相邻节点
	return dfsPerimeter(grid, r-1, c) + dfsPerimeter(grid, r+1, c) + dfsPerimeter(grid, r, c-1) + dfsPerimeter(grid, r, c+1)
}

func islandPerimeter(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return dfsPerimeter(grid, i, j)
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
}
