// [L48-中等] 旋转图像
package main

// 借助新数组
func rotate(matrix [][]int) {
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

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
