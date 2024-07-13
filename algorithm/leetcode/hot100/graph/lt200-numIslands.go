// [L463-中等] 岛屿的周长
package main

import "fmt"

func dfsIslands(grid [][]byte, r, c int) {
	// 超出边界，返回
	if r < 0 || len(grid) <= r || c < 0 || len(grid[0]) <= c {
		return
	}
	// 如果这个格子不是岛屿，直接返回
	if grid[r][c] != '1' {
		return
	}
	// 将遍历过的格子标记为2
	grid[r][c] = '2'
	// 访问上、下、左、右相邻节点
	dfsIslands(grid, r-1, c)
	dfsIslands(grid, r+1, c)
	dfsIslands(grid, r, c-1)
	dfsIslands(grid, r, c+1)
}

func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				dfsIslands(grid, i, j)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
}
