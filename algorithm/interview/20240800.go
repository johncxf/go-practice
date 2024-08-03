package main

import (
	"fmt"
	"strconv"
)

func isStraight(arr []string) bool {
	count := 0
	buckets := make([]int, 11)
	for _, v := range arr {
		if v == "A" {
			buckets[1]++
		} else if v == "大王" || v == "小王" {
			count++
		} else {
			key, _ := strconv.Atoi(v)
			buckets[key]++
		}
	}
	tmp := make([]int, 0)
	for i, v := range buckets {
		j := v
		for j > 0 {
			tmp = append(tmp, i)
			j--
		}
	}
	for i := 1; i < len(tmp); i++ {
		if tmp[i-1]+1 != tmp[i] {
			if count >= tmp[i]-tmp[i-1]-1 {
				count -= tmp[i] - tmp[i-1] - 1
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isStraight([]string{"A", "3", "大王", "小王", "大王"}))
	fmt.Println(isStraight([]string{"A", "2", "5", "小王", "大王"}))
	fmt.Println(isStraight([]string{"A", "2", "6", "小王", "大王"}))

}
