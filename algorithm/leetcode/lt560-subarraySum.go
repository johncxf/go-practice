package main

import "fmt"

// [L560-中等] 和为 K 的子数组
func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	hash := make(map[int]int)
	hash[0] = 1

	for _, num := range nums {
		sum += num
		if val, ok := hash[sum-k]; ok {
			count += val
		}
		hash[sum]++
	}
	return count
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum([]int{-1, -1, 1}, 1))
	fmt.Println(subarraySum([]int{1, -1, 0}, 0))
}
