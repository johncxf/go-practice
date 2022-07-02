package main

import "fmt"

/**
 * 冒泡排序
 *
 * @param input int 整型一维数组
 * @param k int 整型
 * @return int 整型一维数组
 */
func getLeastNumbers(input []int, k int) []int {
	result := []int{}
	if k == 0 {
		return result
	}

	count := len(input)
	if count < k {
		return result
	}

	for i := 0; i < count; i++ {
		for j := 0; j < count-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	result = input[:k]

	return result
}

func quickSort(arr []int, left int, right int) {
	if left > right {
		return
	}

	tmp := arr[left]
	i := left
	j := right

	for i != j {
		for arr[j] >= tmp && i < j {
			j--
		}
		for arr[i] <= tmp && i < j {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left], arr[i] = arr[i], arr[left]

	quickSort(arr, left, i-1)
	quickSort(arr, j+1, right)
}

/**
 * 快排
 *
 * @param input int 整型一维数组
 * @param k int 整型
 * @return int 整型一维数组
 */
func getLeastNumbers2(input []int, k int) []int {
	result := []int{}
	if k == 0 {
		return result
	}

	count := len(input)
	if count < k {
		return result
	}

	quickSort(input, 0, count-1)
	result = input[:k]

	return result
}

func main() {
	ret := getLeastNumbers2([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	fmt.Println(ret)
}
