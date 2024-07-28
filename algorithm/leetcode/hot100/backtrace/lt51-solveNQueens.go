// [L51-困难] N皇后
package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	// 初始化 n*n 大小的棋盘
	state := make([][]string, n)
	for i := range state {
		row := make([]string, n)
		for j := range row {
			row[j] = "."
		}
		state[i] = row
	}
	// 记录列是否放置皇后
	cols := make([]bool, n)
	// 记录对角线是否放置了皇后
	diags1, diags2 := make([]bool, 2*n-1), make([]bool, 2*n-1)
	// 结果集
	res := make([][]string, 0)
	// 进行逐行枚举回溯
	backtraceNQ(n, 0, cols, diags1, diags2, state, &res)
	return res
}

func backtraceNQ(n, row int, cols, diags1, diags2 []bool, state [][]string, res *[][]string) {
	// 放置完所有行，记录结果
	if row == n {
		newState := make([]string, len(state))
		// 将 state 的结果进行组装
		for i := range newState {
			newState[i] = strings.Join(state[i], "")
		}
		// 将当前结果加入结果集
		*res = append(*res, newState)
		return
	}
	// 遍历所有列
	for col := 0; col < n; col++ {
		// 计算两个对角线
		// 对角线 \，这一条对角线上的坐标 row-col 是恒定值，由于存在负数，所以加上 n-1
		diag1 := (row - col) + (n - 1)
		// 对角线 /，这一条对角线上的坐标 row + col 为恒定值
		diag2 := row + col
		// 剪枝：当前列、对角线都没有放置皇后
		if !cols[col] && !diags1[diag1] && !diags2[diag2] {
			// 在当前位置放置皇后
			state[row][col] = "Q"
			// 更新当前列和对角线
			cols[col], diags1[diag1], diags2[diag2] = true, true, true
			// 递归下一行
			backtraceNQ(n, row+1, cols, diags1, diags2, state, res)
			// 回退
			state[row][col] = "."
			cols[col], diags1[diag1], diags2[diag2] = false, false, false
		}
	}
}

func main() {
	fmt.Println(solveNQueens(4))
}
