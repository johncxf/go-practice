package main

import "fmt"

/**
 * [JZ53-简单] 数字在升序数组中出现的次数
 *
 * @param data int 整型一维数组
 * @param k int 整型
 * @return int 整型
 */
func getNumFromAscArray(data []int, k int) int {
	if 0 == len(data) {
		return 0
	}

	// 二分查找 k 位置
	left, right, mid, isExist := 0, len(data)-1, 0, false
	for left <= right {
		mid = (left + right) / 2
		if data[mid] == k {
			isExist = true
			break
		} else if data[mid] > k {
			right = right - 1
		} else {
			left = left + 1
		}
	}

	if false == isExist {
		return 0
	}

	// // 向左遍历查找
	count, move := 1, mid
	for {
		move--
		if move < 0 || data[move] != k {
			break
		}
		count++
	}

	// 向右遍历查找
	move = mid
	for {
		move++
		if move >= len(data) || data[move] != k {
			break
		}
		count++
	}

	return count
}

func main() {
	arr := []int{3, 3, 3, 3}
	res := getNumFromAscArray(arr, 3)
	fmt.Println(res)
}
