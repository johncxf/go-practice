// 链式队列
package main

import (
	"fmt"
)

// QueueNode 定义链表节点
type QueueNode struct {
	Value int
	Next  *QueueNode
}

// 初始化队列
var queueSize = 0
var queue = new(QueueNode)

// Push 入队（从队头插入）
func Push(t *QueueNode, v int) bool {
	if queue == nil {
		queue = &QueueNode{v, nil}
		queueSize++
		return true
	}

	t = &QueueNode{v, nil}
	t.Next = queue
	queue = t
	queueSize++

	return true
}

// Pop 出队（从队尾删除）
func Pop(t *QueueNode) (int, bool) {
	if queueSize == 0 {
		return 0, false
	}

	if queueSize == 1 {
		queue = nil
		queueSize--
		return t.Value, true
	}

	// 迭代队列，直到队尾
	temp := t
	for (t.Next) != nil {
		temp = t
		t = t.Next
	}

	v := (temp.Next).Value
	temp.Next = nil

	queueSize--
	return v, true
}

// 遍历队列所有节点
func traverseQueue(t *QueueNode) {
	if queueSize == 0 {
		fmt.Println("空队列!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	// 入队
	Push(queue, 10)
	fmt.Println("Size:", queueSize)
	// 遍历
	traverseQueue(queue)

	// 出队
	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queueSize)

	// 批量入队
	for i := 0; i < 5; i++ {
		Push(queue, i)
	}
	// 再次遍历
	traverseQueue(queue)
	fmt.Println("Size:", queueSize)

	// 出队
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queueSize)

	// 再次出队
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queueSize)
	// 再次遍历
	traverseQueue(queue)
}
