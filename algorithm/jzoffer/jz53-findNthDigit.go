package main

import (
	"fmt"
	"strconv"
)

/**
 * [JZ44-简单] 数字序列中某一位的数字
 *
 * @param data int 整型一维数组
 * @param k int 整型
 * @return int 整型
 */
func findNthDigit(n int) int {
	// 位数、区间首位数字、区间总位数
	digit, start, sum := 1, 1, 9

	// 获取位数、区间首位、区间总位数
	for n > sum {
		n -= sum
		start *= 10
		digit += 1
		sum = 9 * start * digit
	}

	// 定位 n 在哪个数上
	num := start + (n-1)/digit

	// 定位 n 在这个数的哪一位上
	index := (n - 1) % digit

	// 将 num 转化成 string
	numStr := strconv.Itoa(num)

	// 根据 index 索引获取字符串值
	numStr = string(numStr[index])

	// 将字符串转化成 int
	ret, _ := strconv.Atoi(numStr)

	return ret
}

func main() {
	res := findNthDigit(13)
	fmt.Println(res)
}
