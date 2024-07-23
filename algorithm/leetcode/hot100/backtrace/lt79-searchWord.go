// [L79-中等] 单词搜索
package main

import "fmt"

func searchWord(board [][]byte, word string) bool {
	for i, row := range board {
		for j, col := range row {
			if col == word[0] {
				if backtrackSW(board, word, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

// 回溯
func backtrackSW(board [][]byte, word string, i, j, k int) bool {
	// 越界
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	// 已访问过退出
	if board[i][j] == '0' {
		return false
	}
	// 当前字符串不匹配
	if board[i][j] != word[k] {
		return false
	}
	// 全部匹配通过
	if k == len(word)-1 {
		return true
	}
	// 标记访问过的路径
	board[i][j] = '0'
	// DFS遍历所有路径
	res := backtrackSW(board, word, i-1, j, k+1) ||
		backtrackSW(board, word, i+1, j, k+1) ||
		backtrackSW(board, word, i, j-1, k+1) ||
		backtrackSW(board, word, i, j+1, k+1)
	// 回退
	board[i][j] = word[k]
	return res
}

func main() {
	fmt.Println(searchWord([][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}, "AAB"))
}
