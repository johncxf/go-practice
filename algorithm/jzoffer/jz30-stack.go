package main

import (
	"fmt"
)

var stack []int

func push(node int) {
	stack = append(stack, node)
}

func pop() {
	stack = stack[0 : len(stack)-1]
}

func top() int {
	result := stack[len(stack)-1]
	return result
}

func min() int {
	min := stack[0]
	for _, v := range stack {
		if min > v {
			min = v
		}
	}
	return min
}

func main() {
	push(-1)
	push(2)

	a := min()
	b := top()
	fmt.Println(a)
	fmt.Println(b)
}
