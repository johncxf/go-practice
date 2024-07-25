// [L142-中等] 环形链表 II
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

func detectCycle(head *ListNode) *ListNode {
	hash := map[*ListNode]bool{}
	for head != nil {
		if _, ok := hash[head]; ok {
			return head
		}
		hash[head] = true
		head = head.Next
	}
	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		// 慢指针一步步移动
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		// 快指针2步移动
		fast = fast.Next.Next
		// 当快慢指针相遇
		if fast == slow {
			// cur 指针从头开始移动，slow继续移动，会在入环点相遇
			cur := head
			for cur != slow {
				cur = cur.Next
				slow = slow.Next
			}
			return cur
		}
	}
	return nil
}

func main() {
	var head1 *ListNode
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 0)
	head1 = AddNode(head1, -4)
	fmt.Println(detectCycle(head1))
}
