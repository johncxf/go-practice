package main

import (
	. "go_practice/pkg/datastruct"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, cur *ListNode
	add := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + add
		sum, add = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			cur = head
		} else {
			cur.Next = &ListNode{Val: sum}
			cur = cur.Next
		}
	}

	if add > 0 {
		cur.Next = &ListNode{Val: add}
	}

	return head
}

func main() {
	var head1, head2 *ListNode
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 3)

	head2 = AddNode(head2, 5)
	head2 = AddNode(head2, 6)
	head2 = AddNode(head2, 4)

	rHead := addTwoNumbers(head1, head2)
	TraverseSingleList(rHead)
}
