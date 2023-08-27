package datastruct

import "fmt"

// ListNode 单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// Head 初始化头结点
var Head = new(ListNode)

// AddNode 添加节点
func AddNode(t *ListNode, v int) int {
	if Head == nil {
		t = &ListNode{v, nil}
		Head = t
		return 0
	}
	// 如果节点已经存在
	//if v == t.Val {
	//	fmt.Println("节点已存在:", v)
	//	return -1
	//}
	// 如果当前节点下一个节点为空
	if t.Next == nil {
		t.Next = &ListNode{v, nil}
		return -2
	}
	// 如果当前节点下一个节点不为空
	return AddNode(t.Next, v)
}

// AddNode2 添加节点
func AddNode2(head *ListNode, v int) *ListNode {
	newNode := &ListNode{Val: v, Next: nil}

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

// SearchSingleListNode 查找单链表节点
func SearchSingleListNode(t *ListNode, v int) bool {
	if Head == nil {
		t = &ListNode{v, nil}
		Head = t
		return false
	}
	if v == t.Val {
		return true
	}
	if t.Next == nil {
		return false
	}
	return SearchSingleListNode(t.Next, v)
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
