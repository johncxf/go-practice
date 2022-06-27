package main

import (
	"fmt"
	"strconv"
)

/**
 * [JZ45-中等] 把数组排成最小的数
 *
 * @param numbers int整型一维数组
 * @return string字符串
 */
func printMinNumFromArray(numbers []int) string {
	count := len(numbers)
	if count == 0 {
		return ""
	}

	for i := 0; i < count; i++ {
		for j := 0; j < count-1; j++ {
			str1 := strconv.Itoa(numbers[j]) + strconv.Itoa(numbers[j+1])
			str2 := strconv.Itoa(numbers[j+1]) + strconv.Itoa(numbers[j])
			if str1 > str2 {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	result := ""
	for _, v := range numbers {
		result += strconv.Itoa(v)
	}

	return result
}

func main() {
	ret := printMinNumFromArray([]int{3, 32, 321})
	fmt.Println(ret)
}
