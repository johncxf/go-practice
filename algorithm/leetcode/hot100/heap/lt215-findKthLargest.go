// [L215-中等] 数组中的第K个最大元素
package main

import "fmt"

// 基于堆排序
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	// 构建大顶堆
	for i := n / 2; i >= 0; i-- {
		maxHeapify(nums, i, n)
	}
	// 不断删除堆顶元素，只需要删除k个元素就好了
	for i := len(nums) - 1; i > n-k; i-- {
		// 通过把堆顶和堆底元素进行交换，再删除堆底元素实现删除
		nums[0], nums[i] = nums[i], nums[0]
		// 删除后重新进行堆化
		maxHeapify(nums, 0, i)
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

// 基于快速排序
func findKthLargest2(nums []int, k int) int {
	n := len(nums)
	return quickSort(nums, 0, n-1, k-1)
}

func quickSort(nums []int, left, right, k int) int {
	if left == right {
		return nums[k]
	}
	base := nums[left]
	i, j := left, right
	for i < j {
		// j 向左移动，找到大于 base 的数
		for i < j && nums[j] <= base {
			j--
		}
		// i 向右移动，找到小于 base 的数
		for i < j && nums[i] >= base {
			i++
		}
		// 交换i、j位置
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 交换基准位置
	nums[left], nums[j] = nums[j], nums[left]
	if k <= i {
		return quickSort(nums, left, i, k)
	} else {
		return quickSort(nums, i+1, right, k)
	}
}

// 基于桶排序
func findKthLargest3(nums []int, k int) int {
	// 因为：-10^4 <= nums[i] <= 10^4
	// 所以分配 20001 个桶
	buckets := make([]int, 20001)
	// 将元素放入桶中
	for _, num := range nums {
		// 将负数转为正数
		buckets[num+10000]++
	}
	// 遍历桶
	for i := len(buckets) - 1; i >= 0; i-- {
		k = k - buckets[i]
		if k <= 0 {
			// 加入数组的时候+10000，这里减去
			return i - 10000
		}
	}
	return 0
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))

	fmt.Println(findKthLargest2([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest2([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))

	fmt.Println(findKthLargest3([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest3([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
