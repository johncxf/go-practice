// [L48-中等] 旋转图像
package main

// 借助新数组
// - 时间复杂度：O(N^2)
// - 空间复杂度：O(N^2)
func rotate1(matrix [][]int) {
	n := len(matrix)
	newMatrix := make([][]int, n)
	for i := range newMatrix {
		newMatrix[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, col := range row {
			newMatrix[j][n-1-i] = col
		}
	}
	copy(matrix, newMatrix)
	//fmt.Println(matrix)
}

// 先水平翻转，再对角线翻转
// - 时间复杂度：O(N^2)
// - 空间复杂度：O(1)
func rotate2(matrix [][]int) {
	n := len(matrix)
	// 先水平翻转（直接一行一行互换即可）
	for row := 0; row < n/2; row++ {
		matrix[row], matrix[n-1-row] = matrix[n-1-row], matrix[row]
	}
	// 再主对角线翻转
	for row := 0; row < n; row++ {
		// 这里只需要遍历对角线左侧的元素，和右侧进行换值
		for col := 0; col < row; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}
	//fmt.Println(matrix)
}

func main() {
	rotate1([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	rotate2([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
