// [L73-中等] 矩阵置零
package main

func setZeroes(matrix [][]int) {
	// 初始化两个数组，存放横纵每个位置是否为0
	row, col := make([]bool, len(matrix)), make([]bool, len(matrix[0]))
	// 遍历矩阵，进行标记
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	// 遍历矩阵，对标记位置进行更新
	for i, v := range matrix {
		for j := range v {
			if row[i] || col[j] {
				v[j] = 0
			}
		}
	}
	//fmt.Println(matrix)
}

func main() {
	setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
}
