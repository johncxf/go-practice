package main

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	// 第一次从右侧开始遍历，找出最靠近右侧的较小值
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		// 第二次遍历，找出最靠右的大于较小值的第一个数
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		// 交换两个数的位置
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 将右侧翻转
	reverse(nums[i+1:])
	//fmt.Println(nums)
}

// 翻转数组
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func main() {
	nextPermutation([]int{4, 5, 2, 6, 3, 1})
}
