package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
	"reflect"
)

/**
 * [JZ52-简单] 两个链表的第一个公共结点
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func findFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	p1, p2 := pHead1, pHead2
	for reflect.DeepEqual(p1, p2) == false {
		if p1 == p2 {
			fmt.Println(p1.Val)
		}
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = pHead2
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = pHead1
		}
	}
	return p1
}

func main() {
	fmt.Println("")
	var head1 *ListNode
	var head2 *ListNode

	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 6)
	head1 = AddNode(head1, 7)

	head2 = AddNode(head2, 4)
	head2 = AddNode(head2, 5)
	head2 = AddNode(head2, 6)
	head2 = AddNode(head2, 7)

	rHead := findFirstCommonNode(head1, head2)
	TraverseSingleList(rHead)
}
