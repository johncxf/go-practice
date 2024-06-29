package main

import "fmt"

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 腐烂橘子坐标
	queue := make([][]int, 0)
	// 新鲜橘子的数量
	count := 0
	// 遍历网格，获取新鲜橘子数量和腐烂橘子坐标集合
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	// 腐烂遍历次数 == 分钟
	round := 0
	for count > 0 && len(queue) > 0 {
		round++
		// 获取当前队列长度
		size := len(queue)
		for i := 0; i < size; i++ { // 使用size来避免在迭代中改变queue的长度
			orange := queue[0]
			queue = queue[1:]
			r, c := orange[0], orange[1]
			for r-1 >= 0 && grid[r-1][c] == 1 { // 上
				grid[r-1][c] = 2
				count--
				queue = append(queue, []int{r - 1, c})
			}
			if r+1 < m && grid[r+1][c] == 1 { // 下
				grid[r+1][c] = 2
				count--
				queue = append(queue, []int{r + 1, c})
			}
			if c-1 >= 0 && grid[r][c-1] == 1 { // 左
				grid[r][c-1] = 2
				count--
				queue = append(queue, []int{r, c - 1})
			}
			if c+1 < n && grid[r][c+1] == 1 { // 右
				grid[r][c+1] = 2
				count--
				queue = append(queue, []int{r, c + 1})
			}
		}
	}
	if count > 0 {
		return -1
	}
	return round
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}
