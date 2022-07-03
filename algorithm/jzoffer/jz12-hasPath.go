package main

import (
	"fmt"
)

func dfsMatrix(matrix [][]byte, n int, m int, i int, j int, word string, k int, flag [][]bool) bool {
	// i, j 越界，或字符不匹配、或该点已走过则退出
	if i < 0 || i >= n || j < 0 || j >= m || (matrix[i][j] != word[k]) || flag[i][j] == true {
		return false
	}

	// 已经找到
	if k == len(word)-1 {
		return true
	}

	// 标记当前位置已经走过
	flag[i][j] = true

	// 向四个方向查找
	if dfsMatrix(matrix, n, m, i-1, j, word, k+1, flag) == true ||
		dfsMatrix(matrix, n, m, i+1, j, word, k+1, flag) == true ||
		dfsMatrix(matrix, n, m, i, j-1, word, k+1, flag) == true ||
		dfsMatrix(matrix, n, m, i, j+1, word, k+1, flag) == true {
		return true
	}

	flag[i][j] = false
	return false
}

/**
 * [JZ12-中等] 矩阵中的路径
 *
 * @param matrix char 字符型二维数组
 * @param word string 字符串
 * @return bool 布尔型
 */
func hasPath(matrix [][]byte, word string) bool {
	if len(matrix) == 0 {
		return false
	}

	// 矩阵 n、m
	n, m := len(matrix), len(matrix[0])

	// 初始化一个标记矩阵，记录矩阵走过路径
	flag := make([][]bool, n)
	for k := 0; k < n; k++ {
		flag[k] = make([]bool, m)
	}

	// 遍历矩阵，查找以每个点为起始点情况下是否存在该路径
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 查找路径
			if dfsMatrix(matrix, n, m, i, j, word, 0, flag) == true {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(hasPath([][]byte{{'a', 'b', 'c', 'e'}, {'s', 'f', 'c', 's'}, {'a', 'd', 'e', 'e'}}, "abcced"))
}
