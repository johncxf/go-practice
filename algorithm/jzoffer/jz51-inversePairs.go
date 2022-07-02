package main

import (
	"fmt"
)

var count int = 0

/**
 * 归并排序法
 *
 * @param data int 整型一维数组
 * @return int 整型
 */
func inversePairs(data []int) int {
	n := len(data)
	if n < 2 {
		return 0
	}

	mergeSort(data, 0, len(data)-1)

	return count
}

func mergeSort(arr []int, left int, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2

	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)

	merge(arr, left, mid, right)
}

func merge(arr []int, left int, mid int, right int) {
	result := make([]int, right-left+1)
	c, s, l, r := 0, left, left, mid+1

	for l <= mid && r <= right {
		if arr[l] <= arr[r] {
			result[c] = arr[l]
			l++
		} else {
			count += mid + 1 - l
			count %= 1000000007

			result[c] = arr[r]
			r++
		}
		c++
	}

	for l <= mid {
		result[c] = arr[l]
		c++
		l++
	}

	for r <= right {
		result[c] = arr[r]
		c++
		r++
	}

	for i := 0; i < len(result); i++ {
		arr[s] = result[i]
		s++
	}
}

func main() {
	ret := inversePairs([]int{1, 2, 3, 4, 5, 6, 7, 0})
	fmt.Println(ret)
}
