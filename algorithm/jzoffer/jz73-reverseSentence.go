package main

import (
	"fmt"
	"strings"
)

/**
 * [JZ73-简单] 翻转单词序列
 *
 *
 * @param str string 字符串
 * @return string 字符串
 */
func reverseSentence(str string) string {
	strArr := strings.Split(str, " ")
	result := ""
	for i := len(strArr) - 1; i >= 0; i-- {
		result += strArr[i]
		if i > 0 {
			result += " "
		}
	}
	return result
}

func main() {
	str := "nowcoder. a am I"
	str = reverseSentence(str)
	fmt.Println(str)
}
