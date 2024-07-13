// [L279-中等] 完全平方数
package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			if f[i-j*j] < min {
				min = f[i-j*j]
			}
		}
		f[i] = min + 1
	}
	return f[n]
}

func main() {
	fmt.Println(numSquares(12))
}
