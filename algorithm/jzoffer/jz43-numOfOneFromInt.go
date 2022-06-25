package main

import "fmt"

/**
 * [JZ43-中等] 从 1 到 n 整数中 1 出现的次数
 *
 * @param n int整型
 * @return int整型
 */
func numOfOneFromInt(n int) int {
	i, count, high, cur, low := 1, 0, 0, 0, 0
	for i <= n {
		// 取高位
		high = n / (i * 10)
		// 取当前位
		cur = (n / i) % 10
		// 取低位
		low = n - n/i*i
		if cur == 0 {
			count += high * i
		} else if cur == 1 {
			count += high*i + low + 1
		} else {
			count += (high + 1) * i
		}
		i *= 10
	}
	return count
}

func main() {
	ret := numOfOneFromInt(12013)
	fmt.Println(ret)
}
