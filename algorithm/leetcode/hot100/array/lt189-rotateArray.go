// [L189-中等] 轮转数组
package main

// 切分数组
func rotateArray1(nums []int, k int) {
	n := len(nums)
	// 处理 k > n 的情况
	k = k % n
	arr1 := nums[:n-k]
	arr2 := nums[n-k:]
	ans := append(arr2, arr1...)
	copy(nums, ans)
}

// 翻转数组
func rotateArray2(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

// 反转数组
func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	rotateArray1([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotateArray1([]int{-1, -100, 3, 99}, 2)
	rotateArray1([]int{-1}, 2)

	rotateArray2([]int{1, 2, 3, 4, 5, 6, 7}, 2)
}
