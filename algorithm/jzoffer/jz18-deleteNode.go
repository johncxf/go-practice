package main

import (
	. "go_practice/pkg/datastruct"
)

/**
* [JZ18-简单] 删除链表的节点
*
* @param head ListNode类
* @param val int整型
* @return ListNode类
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	pre := head
	for nil != head {
		if head.Val == val {
			head.Val = head.Next.Val
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}
	return pre
}

func main() {
	var head1 *ListNode
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 6)
	head1 = AddNode(head1, 7)

	rHead := deleteNode(head1, 3)
	TraverseSingleList(rHead)
}
