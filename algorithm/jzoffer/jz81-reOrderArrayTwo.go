package main

import "fmt"

/**
 * [JZ81-简单] 调整数组顺序使奇数位于偶数前面(二)
 *
 * @param array int 整型一维数组
 * @return int 整型一维数组
 */
func reOrderArrayTwo(array []int) []int {
	n := len(array)
	if n <= 1 {
		return array
	}

	var arr []int
	for _, v := range array {
		if v%2 == 1 {
			arr = append([]int{v}, arr...)
		} else {
			arr = append(arr, v)
		}
	}
	return arr
}

/**
 * [JZ81-简单] 调整数组顺序使奇数位于偶数前面(二)
 *
 * @param array int 整型一维数组
 * @return int 整型一维数组
 */
func reOrderArrayTwo2(array []int) []int {
	n := len(array)
	if n <= 1 {
		return array
	}

	odd := 0
	for _, v := range array {
		if v%2 == 1 {
			odd++
		}
	}

	left, right := 0, odd
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if array[i]%2 == 1 {
			arr[left] = array[i]
			left++
		} else {
			arr[right] = array[i]
			right++
		}
	}
	return arr
}

func main() {
	ret := reOrderArrayTwo2([]int{1, 2, 3, 4})
	fmt.Println(ret)
}
