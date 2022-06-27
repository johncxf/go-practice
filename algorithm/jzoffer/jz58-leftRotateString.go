package main

import "fmt"

/**
 * [JZ58-中等] 左旋转字符串
 *
 * @param str string 字符串
 * @param n int 整型
 * @return string 字符串
 */
func leftRotateString(str string, n int) string {
	if 0 == len(str) {
		return ""
	}

	n = n % len(str)

	return str[n:] + str[:n]
}

func main() {
	str := "abcXYZdef"
	ret := leftRotateString(str, 3)
	fmt.Println(ret)
}
