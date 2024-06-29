// [L695-中等] 岛屿的最大面积
package main

import "fmt"

// 判断格子是否在边界内
func inArea(grid [][]int, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}

// 深度优先搜索求面积
func dfsArea(grid [][]int, r, c int) int {
	// 超出边界，返回
	if !inArea(grid, r, c) {
		return 0
	}
	// 如果这个格子不是岛屿，直接返回
	if grid[r][c] != 1 {
		return 0
	}
	// 将遍历过的格子标记为2
	grid[r][c] = 2
	// 访问上、下、左、右相邻节点
	return 1 +
		dfsArea(grid, r-1, c) +
		dfsArea(grid, r+1, c) +
		dfsArea(grid, r, c-1) +
		dfsArea(grid, r, c+1)
}

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := dfsArea(grid, i, j)
				if area > max {
					max = area
				}
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}))
}
