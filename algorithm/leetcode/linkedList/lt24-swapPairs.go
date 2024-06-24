package main

import (
	. "go_practice/pkg/datastruct"
)

// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// 迭代
func swapPairs2(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next

		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		temp = node1
	}
	return dummyHead.Next
}

func main() {
	var (
		head1 *ListNode
		head2 *ListNode
	)
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 4)

	head2 = AddNode(head2, 1)

	TraverseSingleList(swapPairs2(head1))
	TraverseSingleList(swapPairs2(head2))
}
