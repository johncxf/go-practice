// [L300-中等] 最长递增子序列
package main

import "fmt"

func lengthOfLIS(nums []int) int {
	n := len(nums)
	// 构建dp表
	// dp[i] 代表以位置 i 结尾的最长递增子序列的长度
	dp := make([]int, n+1)
	dp[0] = 1
	// dp中的最大值
	max := dp[0]
	// 状态转移，求解 dp[i]
	for i := 1; i < n; i++ {
		// 最小就是本身，为1
		dp[i] = 1
		for j := 0; j < i; j++ {
			// 说明 nums[i] 可以追加到 dp[j] 中
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				} else {
					dp[i] = dp[i]
				}
			}
		}
		// 更新最大值
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// 【扩展】打印最长递增子序列
func resultOfLIS(nums []int) []int {
	n := len(nums)
	result := make([][]int, 0)
	// 第一项默认为第一个元素
	result = append(result, []int{nums[0]})
	// 不断追加元素
	for i := 1; i < n; i++ {
		num := nums[i]
		// 遍历 result 所有序列，将 num 更新到序列中
		isUpdate := false
		for j := len(result) - 1; j >= 0; j-- {
			tmp := result[j]
			tail := tmp[len(tmp)-1]
			// 当前元素大于该序列最后一个元素，则追加一个元素，生成新的序列插入result
			if num > tail {
				tmp = append(tmp, num)
				result = append(result, tmp)
				isUpdate = true
				break
			}
		}
		// 没有能更新进入，说明 num 比 result 中元素都要小，则直接更新 result[0]
		if !isUpdate {
			result[0] = []int{num}
		}
	}
	return result[len(result)-1]
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))

	//fmt.Println(resultOfLIS([]int{4, 5, 1, 2, 7, 3, 6, 9}))
	fmt.Println(resultOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))

}
