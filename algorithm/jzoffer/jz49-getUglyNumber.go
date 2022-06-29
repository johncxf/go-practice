package main

import "fmt"

/**
 * [JZ49-中等] 丑数
 *
 * @param index int整型
 * @return int整型
 */
func getUglyNumber(index int) int {
	if index <= 6 {
		return index
	}

	a, b, c := 0, 0, 0

	data := make([]int, index)
	data[0] = 1

	for i := 1; i < index; i++ {
		min := data[a] * 2
		if t := data[b] * 3; t < min {
			min = t
		}
		if t := data[c] * 5; t < min {
			min = t
		}
		data[i] = min
		if min == data[a]*2 {
			a++
		}
		if min == data[b]*3 {
			b++
		}
		if min == data[c]*5 {
			c++
		}
	}
	return data[index-1]
}

func main() {
	ret := getUglyNumber(7)
	fmt.Println(ret)
}
