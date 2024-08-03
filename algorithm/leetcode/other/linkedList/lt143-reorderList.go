package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

// 线性表
// - 时间复杂度：O(N)
// - 空间复杂度：O(N)
func reorderList(head *ListNode) {
	nodes := []*ListNode{}
	// 将链表所有节点存入数组
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		nodes[j].Next = nodes[i+1]
		i++
		j--
	}
	nodes[i].Next = nil
}

// 找链表中点并拆分为两个链表，反转右侧部分，在合并左右部分
// - 时间复杂度：O(N)
// - 空间复杂度：O(1)
func reorderList2(head *ListNode) {
	if head.Next == nil {
		return
	}
	// 中间节点
	mid := middleList(head)
	// 左右链表
	left, right := head, mid.Next
	// 断开左右链表
	mid.Next = nil
	// 反转右侧链表
	right = reverseList(right)
	// 合并链表
	mergeTwoList(left, right)
}

// 合并链表
func mergeTwoList(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		tmp1 := l1.Next
		tmp2 := l2.Next

		l1.Next = l2
		l1 = tmp1

		l2.Next = l1
		l2 = tmp2
	}
}

// 找出链表中间节点
func middleList(head *ListNode) *ListNode {
	slow, fast := head, head
	// 偶数中间值为前一个
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

func main() {
	head1 := NewListNode(1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 5)
	fmt.Println("head1: ")
	TraverseSingleList(head1)

	fmt.Println("重排后: ")
	//reorderList(head1)
	//TraverseSingleList(head1)

	reorderList2(head1)
	TraverseSingleList(head1)
}
