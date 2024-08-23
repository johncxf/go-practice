// [L287-中等] 寻找重复数
package main

import "fmt"

// 快慢指针
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

// 辅助数组
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

// 利用原数组作为辅助数组
func findDuplicate3(nums []int) int {
	// 遍历所有元素
	for _, num := range nums {
		// 如果该位置为负数，说明被修改过，也就是存在重复
		if nums[abs(num)] < 0 {
			return abs(num)
		}
		// 修改以当前元素绝对值为下标位置的值为负数
		nums[abs(num)] = -nums[abs(num)]
	}
	return -1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))

	fmt.Println(findDuplicate2([]int{1, 3, 4, 2, 2}))

	fmt.Println(findDuplicate3([]int{1, 3, 4, 2, 2}))
}
