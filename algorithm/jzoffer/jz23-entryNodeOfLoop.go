package main

import (
	. "go_practice/pkg/datastruct"
)

/**
 * [JZ23-中等] 链表中环的入口结点
 *
 * @param head ListNode类
 * @return ListNode类
 */
func entryNodeOfLoop(pHead *ListNode) *ListNode {
	if nil == pHead || nil == pHead.Next {
		return nil
	}
	fast, slow := pHead, pHead
	for nil != fast && nil != fast.Next.Next {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = pHead
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}

func main() {
	var head1 *ListNode
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 5)

	kHead := reverseList2(head1)

	TraverseSingleList(kHead)
}
