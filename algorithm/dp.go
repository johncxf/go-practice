// 动态规划经典题型
package main

import (
	"fmt"
)

// 0-1 背包问题
// 题目描述：
// 给定 n 个物品，第 i 个物品的重量为 wgt[i-1]、价值为 val[i-1]，和一个容量为 cap 的背包。
// 每个物品 "只能选择一次"，问在限定背包容量下能放入物品的最大价值。
func knapsackDP(wgt, val []int, cap int) int {
	n := len(wgt)
	// 构建 dp 表
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, cap+1)
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		for c := 1; c <= cap; c++ {
			if wgt[i-1] > c {
				// 若超过背包容量，则不选物品 i
				dp[i][c] = dp[i-1][c]
			} else {
				// 不选和选物品 i 这两种方案的较大值
				if dp[i-1][c] > dp[i-1][c-wgt[i-1]]+val[i-1] {
					dp[i][c] = dp[i-1][c]
				} else {
					dp[i][c] = dp[i-1][c-wgt[i-1]] + val[i-1]
				}
			}
		}
	}
	return dp[n][cap]
}

// 0-1 背包问题 - 空间优化
func knapsackDP2(wgt, val []int, cap int) int {
	n := len(wgt)
	// 构建 dp 表
	dp := make([]int, cap+1)
	// 状态转移
	for i := 1; i <= n; i++ {
		// 改成倒叙遍历
		for c := cap; c >= 1; c-- {
			// wgt[i-1] > c 的情况 dp[c] = dp[c]，可以省略
			if wgt[i-1] <= c {
				// 不选和选物品 i 这两种方案的较大值
				if dp[c] < dp[c-wgt[i-1]]+val[i-1] {
					dp[c] = dp[c-wgt[i-1]] + val[i-1]
				}
			}
		}
	}
	return dp[cap]
}

// 完全背包问题
// 题目描述：
// 给定 n 个物品，第 i 个物品的重量为 wgt[i-1]、价值为 val[i-1]，和一个容量为 cap 的背包。
// 每个物品 "可以重复选取"，问在限定背包容量下能放入物品的最大价值。
func unboundedKnapsackDP(wgt, val []int, cap int) int {
	n := len(wgt)
	// 建立 dp 表
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, cap+1)
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		for c := 1; c <= cap; c++ {
			if wgt[i-1] > c {
				dp[i][c] = dp[i-1][c]
			} else {
				if dp[i-1][c] > dp[i][c-wgt[i-1]]+val[i-1] {
					dp[i][c] = dp[i-1][c]
				} else {
					dp[i][c] = dp[i][c-wgt[i-1]] + val[i-1]
				}
			}
		}
	}
	return dp[n][cap]
}

// 完全背包问题 - 空间优化
func unboundedKnapsackDP2(wgt, val []int, cap int) int {
	n := len(wgt)
	// 建立 dp 表
	dp := make([]int, cap+1)
	// 状态转移
	for i := 1; i <= n; i++ {
		for c := 1; c <= cap; c++ {
			if wgt[i-1] <= c {
				if dp[c] < dp[c-wgt[i-1]]+val[i-1] {
					dp[c] = dp[c-wgt[i-1]] + val[i-1]
				}
			}
		}
	}
	return dp[cap]
}

func main() {
	//0-1 背包问题
	fmt.Println(knapsackDP([]int{1, 2, 3}, []int{5, 11, 15}, 4))
	fmt.Println(knapsackDP2([]int{1, 2, 3}, []int{5, 11, 15}, 4))

	// 完全背包问题
	fmt.Println(unboundedKnapsackDP([]int{1, 2, 3}, []int{5, 11, 15}, 4))
	fmt.Println(unboundedKnapsackDP2([]int{1, 2, 3}, []int{5, 11, 15}, 4))
}
