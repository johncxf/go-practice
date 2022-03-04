// 排序算法
package main

// 冒泡排序
func bubbleSort(arr []int) []int {
	count := len(arr)
	if 1 >= count {
		return arr
	}
	for i := 0; i < count; i++ {
		for j := 0; j < count-1; j++ {
			// 从小到大
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			// 从大到小
			// if arr[j] < arr[j+1] {
			// 	arr[j], arr[j+1] = arr[j+1], arr[j]
			// }
		}
	}
	return arr
}

// 快速排序
func quickSort(arr []int, left int, right int) []int {

}

func main() {
	arr := []int{4, 5, 6, 7, 8, 3, 2, 1}
	arr = bubbleSort(arr)
	// fmt.Println(arr)
}
