package main

import "fmt"

/**
 * 【暴力】数组中重复的数字
 *
 * @param numbers int 整型一维数组
 * @return int 整型
 */
func duplicate(numbers []int) int {
	// 遍历数组
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				return numbers[i]
			}
		}
	}
	return -1
}

/**
 * 位置重排
 *
 * @param numbers int 整型一维数组
 * @return int 整型
 */
func duplicate2(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] != i {
			// 当前值对应的下标已经有正确的对应数值，说明重复
			if numbers[i] == numbers[numbers[i]] {
				return numbers[i]
			}
			// 未存在，则交换位置
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}
	return -1
}

func main() {
	fmt.Println(duplicate2([]int{2, 3, 1, 0, 2, 5, 3}))
}
