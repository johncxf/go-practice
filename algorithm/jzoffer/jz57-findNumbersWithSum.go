package main

import "fmt"

/**
 * 双指针
 *
 * @param array int 整型一维数组
 * @param sum int 整型
 * @return int 整型一维数组
 */
func findNumbersWithSum(array []int, sum int) []int {
	result := []int{}
	l, r := 0, len(array)-1
	for l < r {
		tmpSum := array[l] + array[r]
		if tmpSum == sum {
			result = []int{array[l], array[r]}
			break
		} else if tmpSum < sum {
			l++
		} else {
			r--
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 4, 7, 11, 15}
	ret := findNumbersWithSum(arr, 15)
	fmt.Println(ret)
}
