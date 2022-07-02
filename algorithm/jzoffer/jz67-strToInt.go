package main

import (
	"fmt"
	"math"
	"strings"
)

/**
 * [JZ67-中等] 把字符串转换成整数(atoi)
 *
 * @param s string字符串
 * @return int整型
 */
func strToInt(s string) int {
	// 去掉前后空格
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}

	// 处理符号位
	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	// 判断整数有效部分
	res := 0
	for _, v := range s {
		// 数据范围
		if v >= '0' && v <= '9' {
			// 将符合条件的整数字符串转化后加入的结果整数中
			res = res*10 + int(v-'0')
		} else {
			break
		}
		if res > int(math.Pow(2, 31)) {
			break
		}
	}
	// 乘以符号位
	res = int(sign) * res

	// 数据范围判断
	if res > int(math.Pow(2, 31)-1) {
		return int(math.Pow(2, 31) - 1)
	}
	if res < int(0-math.Pow(2, 31)) {
		return int(0 - math.Pow(2, 31))
	}

	return int(res)
}

func main() {
	ret := strToInt("4396 clearlove")
	fmt.Println(ret)
}
