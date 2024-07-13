package main

import "fmt"

// [L75-中等] 颜色分类
func sortColors(nums []int) []int {
	// p0 为0元素的末尾，p2 为2元素的开始
	p0, p2 := 0, len(nums)-1
	// 当前位置
	i := 0
	for i <= p2 {
		// 为0，则和p0交换位置
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
		} else if nums[i] == 2 { // 为2，则和 p2 交换位置
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		} else { // 为1，则不交换，继续移动
			i++
		}
	}
	return nums
}

func main() {
	fmt.Println(sortColors([]int{2, 0, 2, 1, 1, 0}))
	fmt.Println(sortColors([]int{2, 0, 1}))
}
