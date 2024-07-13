// [L560-中等] 和为 K 的子数组
package main

import (
	"fmt"
)

// 枚举
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func subarraySum1(nums []int, k int) int {
	n := len(nums)
	ans := 0
	for i, num := range nums {
		target := k - num
		if target == 0 {
			ans++
		}
		for j := i + 1; j < n; j++ {
			target -= nums[j]
			if target == 0 {
				ans++
			}
		}
	}
	return ans
}

// 前缀和
// 时间复杂度：O(N^2)
// 空间复杂度：O(N)
func subarraySum2(nums []int, k int) int {
	n := len(nums)
	// 计算前缀和数组
	// preSum[i] = 前 i 和元素之和
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	// 在 preSum 中计算两数之差为 k 的数量
	cnt := 0
	// 依次遍历求解
	for i := 0; i < len(preSum); i++ {
		for j := i + 1; j < len(preSum); j++ {
			if preSum[j]-preSum[i] == k {
				cnt++
			}
		}
	}
	// 优化：结合 hashMap
	//hash := map[int]int{}
	//for _, sum := range preSum {
	//	if val, ok := hash[sum-k]; ok {
	//		cnt += val
	//	}
	//	hash[sum]++
	//}

	return cnt
}

// 前缀和
// 时间复杂度：O(2N)
// 空间复杂度：O(2N)
func subarraySum3(nums []int, k int) int {
	n := len(nums)
	// 计算前缀和数组
	// preSum[i] = 前 i 和元素之和
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	// 在 preSum 中计算两数之差为 k 的数量
	cnt := 0
	//优化：结合 hashMap
	hash := map[int]int{}
	for _, sum := range preSum {
		if val, ok := hash[sum-k]; ok {
			cnt += val
		}
		hash[sum]++
	}

	return cnt
}

// 前缀和+哈希表
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func subarraySum4(nums []int, k int) int {
	count, sum := 0, 0
	hash := make(map[int]int)
	hash[0] = 1

	for _, num := range nums {
		// 直接构建前缀和
		sum += num
		// 判断前缀和是否在 hashMap 中
		if val, ok := hash[sum-k]; ok {
			count += val
		}
		// 将前缀和加入 hashMap
		hash[sum]++
	}
	return count
}

func main() {
	fmt.Println("枚举:")
	fmt.Println(subarraySum1([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum1([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum1([]int{-1, -1, 1}, 1))
	fmt.Println(subarraySum1([]int{1, -1, 0}, 0))

	fmt.Println("前缀和:")
	fmt.Println(subarraySum2([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum2([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum2([]int{-1, -1, 1}, 1))
	fmt.Println(subarraySum2([]int{1, -1, 0}, 0))

	fmt.Println("前缀和+哈希表:")
	fmt.Println(subarraySum3([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum3([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum3([]int{-1, -1, 1}, 1))
	fmt.Println(subarraySum3([]int{1, -1, 0}, 0))

	fmt.Println("前缀和+哈希表（优化）:")
	fmt.Println(subarraySum4([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum4([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum4([]int{-1, -1, 1}, 1))
	fmt.Println(subarraySum4([]int{1, -1, 0}, 0))
}
