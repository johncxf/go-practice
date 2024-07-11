// [L1-简单] 两数之和
package main

import "fmt"

// 哈希表
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for k, v := range nums {
		if p, ok := hash[target-v]; ok {
			return []int{p, k}
		}
		hash[v] = k
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
