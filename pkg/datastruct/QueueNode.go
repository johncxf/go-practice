package datastruct

import "fmt"

// QueueNode 定义链表节点
type QueueNode struct {
	Value interface{}
	Next  *QueueNode
}

// LinkedQueue 链式队列
type LinkedQueue struct {
	head *QueueNode // 队列头部
	tail *QueueNode // 队列尾部
	size int        // 队列大小
}

// NewLinkedQueue 初始化一个空的链式队列
func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Enqueue 入队操作
func (q *LinkedQueue) Enqueue(value interface{}) {
	newNode := &QueueNode{value, nil}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.Next = newNode
		q.tail = newNode
	}
}

// Dequeue 出队操作
func (q *LinkedQueue) Dequeue() interface{} {
	if q.head == nil {
		return nil
	}
	value := q.head.Value
	q.head = q.head.Next
	if q.head == nil {
		q.tail = nil
	}
	return value
}

// Peek 获取队列头部元素值
func (q *LinkedQueue) Peek() interface{} {
	if q.head == nil {
		return nil
	}
	return q.head.Value
}

// IsEmpty 检查队列是否为空
func (q *LinkedQueue) IsEmpty() bool {
	return q.head == nil
}

// Traverse 遍历
func (q *LinkedQueue) Traverse() {
	if q.head == nil {
		fmt.Println("空队列!")
		return
	}
	for q.head != nil {
		fmt.Printf("%d -> ", q.head.Value)
		q.head = q.head.Next
	}
}
