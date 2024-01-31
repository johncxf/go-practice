package main

import (
	"fmt"
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
	fmt.Println("")
	Head = nil
	AddNode(Head, 1)
	AddNode(Head, 2)
	AddNode(Head, 3)
	AddNode(Head, 4)
	AddNode(Head, 5)

	//kHead := reverseList(Head)
	kHead := reverseList2(Head)

	TraverseSingleList(kHead)
}
