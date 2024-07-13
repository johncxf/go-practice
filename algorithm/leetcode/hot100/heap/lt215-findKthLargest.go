// [L215-中等] 数组中的第K个最大元素
package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	// 构建大顶堆
	for i := n / 2; i >= 0; i-- {
		maxHeapify(nums, i, n)
	}
	// 不断删除堆顶元素
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		// 通过把堆顶和堆底元素进行交换，再删除堆底元素实现删除
		// 这样可以保证堆结构最小调整
		nums[0], nums[i] = nums[i], nums[0]
		n--
		// 删除后重新进行堆化
		maxHeapify(nums, 0, n)
	}
	return nums[0]
}

func maxHeapify(a []int, i, size int) {
	// 左子节点、右子节点
	l, r := 2*i+1, 2*i+2
	// 最大节点
	largest := i
	// 左子节点存在，且大于最大节点
	if l < size && a[l] > a[largest] {
		largest = l
	}
	// 右子节点存在，且大于最大节点
	if r < size && a[r] > a[largest] {
		largest = r
	}
	// 最大节点不等于当前节点，则交换位置
	//fmt.Println(largest)
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		// 递归
		maxHeapify(a, largest, size)
	}
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
