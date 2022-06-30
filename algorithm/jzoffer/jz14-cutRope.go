package main

import "fmt"

/**
 * 动态规划
 *
 * @param n int 整型
 * @return int 整型
 */
func cutRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	dp[4] = 4
	for i := 5; i <= n; i++ {
		for j := 1; j < i; j++ {
			if dp[i] > j*dp[i-j] {
				dp[i] = dp[i]
			} else {
				dp[i] = j * dp[i-j]
			}
		}
	}
	return dp[n]
}

func main() {
	ret := cutRope(8)
	fmt.Println(ret)
}
