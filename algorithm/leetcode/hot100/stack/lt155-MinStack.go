// [L155-中等] 最小栈
package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (t *MinStack) Push(val int) {
	t.stack = append(t.stack, val)
	top := t.minStack[len(t.minStack)-1]
	if top < val {
		t.minStack = append(t.minStack, top)
	} else {
		t.minStack = append(t.minStack, val)
	}
}

func (t *MinStack) Pop() {
	t.stack = t.stack[:len(t.stack)-1]
	t.minStack = t.minStack[:len(t.minStack)-1]
}

func (t *MinStack) Top() int {
	return t.stack[len(t.stack)-1]
}

func (t *MinStack) GetMin() int {
	return t.minStack[len(t.minStack)-1]
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
