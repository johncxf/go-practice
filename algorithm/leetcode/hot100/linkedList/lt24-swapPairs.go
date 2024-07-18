// [L24-中等] 两两交换链表中的节点
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
	if head == nil || head.Next == nil {
		return head
	}
	// 哑节点：0-1-2-3-4
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next

		cur.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		cur = node1
	}
	return dummy.Next
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
