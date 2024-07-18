// [L189-中等] 轮转数组
package main

// 切法数组
func rotateArray(nums []int, k int) {
	n := len(nums)
	// 处理 k > n 的情况
	k = k % n
	arr1 := nums[:n-k]
	arr2 := nums[n-k:]
	ans := append(arr2, arr1...)
	copy(nums, ans)
}

func main() {
	rotateArray([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotateArray([]int{-1, -100, 3, 99}, 2)
	rotateArray([]int{-1}, 2)
}
