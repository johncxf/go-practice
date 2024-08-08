// 排序算法
package main

import "fmt"

// 冒泡排序
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func bubbleSort(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		// 内循环：将 [0, i] 元素中最大值换到最右边
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 选择排序
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func selectSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

// 插入排序
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		base := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > base {
				// 交换位置
				arr[j+1], arr[j] = arr[j], base
			} else {
				break
			}
		}
	}
}

// 快速排序
// 时间复杂度：O(nlog2n)，最差：O(n^2)
// 空间复杂度：O(n)，最差：O(log2n)
func quickSort(nums []int, left int, right int) {
	if left > right {
		return
	}
	base := nums[left]
	i, j := left, right
	for i != j {
		// 哨兵 j 向左移动，查找小于基准数 base 时停下
		for nums[j] >= base && i < j {
			j--
		}
		// 哨兵 i 向右移动，查找大于基准数 base 时停下
		for nums[i] <= base && i < j {
			i++
		}
		// 交换 i、j 位置
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 将基准数归位
	nums[left], nums[i] = nums[i], nums[left]
	quickSort(nums, left, i-1)
	quickSort(nums, j+1, right)
}

// 归并排序
// 时间复杂度：O(nlog2n)
// 空间复杂度：O(n)
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}

	// 递归的进行数组拆分
	// 左区间：[left, mid]，右区间：[mid, right]
	middle := length / 2
	left := mergeSort(arr[0:middle])
	right := mergeSort(arr[middle:])

	// 递归进行数组合并
	return merge(left, right)
}

// 递归合并数组（从小到大）
func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

// 堆排序
// 时间复杂度：O(nlogn)
// 空间复杂度：O(1)
func heapSort(arr []int) {
	// 构建大顶堆
	// 这里为什么只遍历一半元素？因为数组后面元素都是前面元素的左右子节点，在堆化过程中都能遍历到
	for i := len(arr)/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, len(arr))
	}
	// 不断进行 将堆顶与堆低元素对换，然后重新进行 堆化 操作
	for i := len(arr) - 1; i > 0; i-- {
		// 交换堆顶与堆底元素
		arr[0], arr[i] = arr[i], arr[0]
		// 堆底出堆
		maxHeapify(arr, 0, i)
	}
}

// 堆化-大顶堆
// 将传入的元素堆化至堆顶位置
func maxHeapify(arr []int, i, heapSize int) {
	// 左子节点、右子节点
	l, r := 2*i+1, 2*i+2
	// 最大值下标
	largest := i
	// 当左子节点大于最大值，则更新最大值
	if l < heapSize && arr[l] > arr[largest] {
		largest = l
	}
	// 当右子节点大于最大值，则更新最大值
	if r < heapSize && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		// 交换位置
		arr[i], arr[largest] = arr[largest], arr[i]
		// 递归
		maxHeapify(arr, largest, heapSize)
	}
}

// 桶排序
func bucketSort(nums []int) []int {
	// 桶分配需要根据实际情况考虑
	buckets := make([]int, 11)
	for _, num := range nums {
		buckets[num]++
	}
	var ans []int
	for i, v := range buckets {
		for v > 0 {
			ans = append(ans, i)
			v--
		}
	}
	return ans
}

// 计数排序（适用整数数组）
func countingSort(nums []int) []int {
	// 计算出数组中最大元素
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	// counter[num] 代表 num 的出现次数
	counter := make([]int, max+1)
	for _, num := range nums {
		counter[num]++
	}
	// 遍历 counter
	var ans []int
	for i, v := range counter {
		for v > 0 {
			ans = append(ans, i)
			v--
		}
	}
	return ans
}

func main() {
	fmt.Println("冒泡排序：")
	arr1 := []int{4, 5, 6, 7, 8, 3, 2, 1}
	bubbleSort(arr1)
	fmt.Println(arr1)

	fmt.Println("选择排序：")
	arr2 := []int{4, 5, 6, 7, 8, 3, 2, 1}
	selectSort(arr2)
	fmt.Println(arr2)

	fmt.Println("插入排序：")
	arr3 := []int{4, 1, 3, 1, 5, 2}
	insertSort(arr3)
	fmt.Println(arr3)

	fmt.Println("快速排序：")
	arr4 := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	quickSort(arr4, 0, len(arr4)-1)
	fmt.Println(arr4)

	fmt.Println("归并排序：")
	arr5 := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	fmt.Println(mergeSort(arr5))

	fmt.Println("堆排序：")
	arr6 := []int{4, 5, 6, 7, 8, 3, 2, 1}
	heapSort(arr6)
	fmt.Println(arr6)

	fmt.Println("桶排序：")
	arr7 := []int{4, 5, 6, 7, 8, 3, 2, 1}
	fmt.Println(bucketSort(arr7))

	fmt.Println("计数排序：")
	arr8 := []int{4, 5, 6, 7, 8, 3, 2, 1}
	fmt.Println(countingSort(arr8))
}
