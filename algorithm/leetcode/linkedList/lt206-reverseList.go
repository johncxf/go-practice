// [L206-简单] 反转链表
package main

import (
	. "go_practice/pkg/datastruct"
)

// 迭代
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	var head1 *ListNode
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 5)

	TraverseSingleList(head1)
	rHead := reverseList(head1)
	TraverseSingleList(rHead)
}
