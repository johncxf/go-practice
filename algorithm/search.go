// 查找算法
package main

import "fmt"

// 二分查找 - 针对递增数组
func binarySearch(data []int, search int) int {
	low := 0
	height := len(data) - 1
	for low < height {
		mid := (height + low) / 2
		if data[mid] == search {
			return mid
		} else if data[mid] > search {
			height = mid - 1
		} else {
			low = height + 1
		}
	}
	return -1
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := binarySearch(data, 4)
	fmt.Println(result)
}
