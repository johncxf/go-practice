package main

import (
	"fmt"
	"math/bits"
)

// [L461-简单] 汉明距离
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	fmt.Println(hammingDistance(1, 4))
	fmt.Println(hammingDistance(3, 1))
}
