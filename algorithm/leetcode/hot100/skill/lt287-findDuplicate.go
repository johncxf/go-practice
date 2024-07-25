// [L287-中等] 寻找重复数
package main

import "fmt"

func findDuplicate(nums []int) int {
	// 定义快慢指针
	slow, fast := 0, 0
	// 慢指针每次走1步，快指针每次走2步，找出相遇点
	slow, fast = nums[slow], nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// 重新设定一个指针 cur，从头开始走，当 cur 和 slow 相遇的时候，就是入环点
	cur := 0
	for cur != slow {
		slow = nums[slow]
		cur = nums[cur]
	}
	return cur
}

func findDuplicate2(nums []int) int {
	tmp := make([]int, len(nums))
	for _, num := range nums {
		if tmp[num] > 0 {
			return num
		}
		tmp[num]++
	}
	return -1
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))

	fmt.Println(findDuplicate2([]int{1, 3, 4, 2, 2}))

}
