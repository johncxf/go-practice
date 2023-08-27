package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

/**
* [JZ25-简单] 合并两个排序的链表
*
* @param pHead1 ListNode类
* @param pHead2 ListNode类
* @return ListNode类
 */
func mergeList(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil && pHead2 == nil {
		return nil
	}
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	if pHead1.Val <= pHead2.Val {
		pHead1.Next = mergeList(pHead1.Next, pHead2)
		return pHead1
	}
	if pHead1.Val > pHead2.Val {
		pHead2.Next = mergeList(pHead1, pHead2.Next)
		return pHead2
	}
	return nil
}

func main() {
	fmt.Println("")
	var head1 *ListNode
	var head2 *ListNode

	head1 = AddNode2(head1, 1)
	head1 = AddNode2(head1, 3)
	head1 = AddNode2(head1, 5)

	head2 = AddNode2(head2, 2)
	head2 = AddNode2(head2, 4)
	head2 = AddNode2(head2, 6)

	rHead := mergeList(head1, head2)
	TraverseSingleList(rHead)
}
