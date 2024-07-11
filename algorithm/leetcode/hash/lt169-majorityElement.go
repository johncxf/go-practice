// [L169-简单] 多数元素
package main

import "fmt"

// 哈希
func majorityElement(nums []int) int {
	hash := make(map[int]int)
	for _, num := range nums {
		if _, ok := hash[num]; ok {
			hash[num]++
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

// 哈希2
func majorityElement2(nums []int) int {
	n := len(nums)
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
	}
	half := n / 2
	if n%2 != 0 {
		half = n/2 + 1
	}
	for num, cnt := range hash {
		if cnt >= half {
			return num
		}
	}
	return 0
}

// 优化
func majorityElement3(nums []int) int {
	winner := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == winner {
			count++
		} else if count == 0 {
			winner = nums[i]
			count++
		} else {
			count--
		}
	}
	return winner
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement2([]int{3, 2, 3}))

	fmt.Println(majorityElement3([]int{3, 2, 3}))
}
