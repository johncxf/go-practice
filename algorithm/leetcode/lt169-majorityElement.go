package main

import "fmt"

func majorityElement(nums []int) int {
	hash := map[int]int{}
	for _, num := range nums {
		if _, ok := hash[num]; ok {
			hash[num] = hash[num] + 1
		} else {
			hash[num] = 1
		}
	}
	maxKey := 0
	maxValue := 0
	for k, v := range hash {
		if v > maxValue {
			maxKey = k
			maxValue = v
		}
	}
	return maxKey
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}