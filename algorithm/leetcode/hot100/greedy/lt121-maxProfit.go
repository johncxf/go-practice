package main

import "fmt"

// [L121-简单] 买卖股票的最佳时机
func maxProfit(prices []int) int {
	maxPro, minPro := 0, int(1e5)
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPro {
			minPro = prices[i]
		}
		if prices[i]-minPro > maxPro {
			maxPro = prices[i] - minPro
		}
	}
	return maxPro
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
