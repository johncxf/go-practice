package main

import (
	. "go_practice/pkg/datastruct"
)

// [L19-中等] 删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := dummy, dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
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
	head1 = AddNode(head1, 5)

	head2 = AddNode(head2, 1)

	TraverseSingleList(removeNthFromEnd(head1, 2))
	TraverseSingleList(removeNthFromEnd(head2, 1))
}
