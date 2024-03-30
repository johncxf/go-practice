package main

import "fmt"

func maxProfit(prices []int) int {
	minPro := int(1e5)
	maxPro := 0
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
