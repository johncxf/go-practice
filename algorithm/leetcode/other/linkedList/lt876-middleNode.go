// [L876-简单] 链表的中间节点
package main

import (
	. "go_practice/pkg/datastruct"
)

// 遍历
// - 时间复杂度：O(N)
// - 空间复杂度：O(1)
func middleNode1(head *ListNode) *ListNode {
	tmp := head
	count := 0
	for tmp != nil {
		count++
		tmp = tmp.Next
	}
	for i := 0; i < count/2; i++ {
		head = head.Next
	}
	return head
}

// 快慢指针
// - 时间复杂度：O(N)
// - 空间复杂度：O(1)
func middleNode2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	head1 := NewListNode(1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 5)
	//TraverseSingleList(middleNode1(head1))
	TraverseSingleList(middleNode2(head1))
}
