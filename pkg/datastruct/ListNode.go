package datastruct

import "fmt"

// ListNode 单链表
type ListNode struct {
	Val  int       // 节点值
	Next *ListNode // 指向下一个节点的指针
}

// NewListNode 创建一个新的链表
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// AddNode 添加节点
func AddNode(head *ListNode, v int) *ListNode {
	//newNode := &ListNode{Val: v, Next: nil}
	newNode := NewListNode(v)

	if head == nil {
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	return head
}

// TraverseSingleList 遍历单链表
func TraverseSingleList(t *ListNode) {
	if t == nil {
		fmt.Println("-> 空链表!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Val)
		t = t.Next
	}
	fmt.Println()
}

// GetSingleListSize 获取链表长度
func GetSingleListSize(t *ListNode) int {
	if t == nil {
		fmt.Println("-> 空链表!")
		return 0
	}
	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}
