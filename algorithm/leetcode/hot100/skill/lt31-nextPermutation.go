// [L31-中等] 下一个排列
package main

func nextPermutation(nums []int) {
	n := len(nums)
	// 从右往左遍历，找出第一个较小值 nums[i]
	i := n - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	// i < 0 的时候相当于当前是最大的序列，直接翻转即可，也就是翻转成最小序列
	if i >= 0 {
		// 从右往左遍历，找出第一个大于较小值的数 nums[j]
		j := n - 1
		for ; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		// 交换 nums[i] 和 nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 翻转 nums[i+1:] 部分
	reverse(nums[i+1:])
	//fmt.Println(nums)
}

// 翻转数组
func reverse(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return
}

func main() {
	nextPermutation([]int{4, 5, 2, 6, 3, 1})
}
