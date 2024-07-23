// [L322-中等] 零钱兑换
package main

import "fmt"

// 动态规划-完全背包问题解题思路
func coinChange(coins []int, amount int) int {
	n := len(coins)
	max := amount + 1
	// 构建dp表
	// dp[i][j]：i个硬币，目标金额为j时的最少金币数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, max)
	}
	// 边界确定
	// 金额为0时候，不需要金币，所以：dp[i][0]=0
	// for i := 0; i <= n; i++ {
	//     dp[i][0] = 0
	// }
	// 当无硬币时，无法凑出任意 >0 的目标金额，即是无效解，用大于 amount+1 表示
	for a := 1; a <= amount; a++ {
		dp[0][a] = max
	}
	// 其他结果
	for i := 1; i <= n; i++ {
		for a := 1; a <= amount; a++ {
			//if coins[i-1] > a {
			//	// 超过目标金额，则只能不选
			//	dp[i][a] = dp[i-1][a]
			//} else {
			//	// 不选和选硬币 i 这两种方案的较小值
			//	if dp[i-1][a] < dp[i][a-coins[i-1]]+1 {
			//		dp[i][a] = dp[i-1][a]
			//	} else {
			//		dp[i][a] = dp[i][a-coins[i-1]] + 1
			//	}
			//}
			dp[i][a] = dp[i-1][a]
			if a >= coins[i-1] && dp[i][a] > dp[i][a-coins[i-1]]+1 {
				dp[i][a] = dp[i][a-coins[i-1]] + 1
			}
		}
	}
	// max 这里代表无效解
	if dp[n][amount] != max {
		return dp[n][amount]
	}
	return -1
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
