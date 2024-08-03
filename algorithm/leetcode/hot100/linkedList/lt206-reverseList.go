// [L206-简单] 反转链表
package main

import (
	. "go_practice/pkg/datastruct"
)

// 先遍历链表，存储在数组中，再重新构建链表
func reverseList1(head *ListNode) *ListNode {
	var tmp []int
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	ansHead := &ListNode{Val: 0}
	tmpHead := ansHead
	for i := len(tmp) - 1; i >= 0; i-- {
		tmpHead.Next = &ListNode{Val: tmp[i]}
		tmpHead = tmpHead.Next
	}
	return ansHead.Next
}

// 先遍历链表，存储在数组中，再重新构建链表
func reverseList2(head *ListNode) *ListNode {
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	dummy := &ListNode{Val: 0}
	cur := dummy
	for i := len(nodes) - 1; i >= 0; i-- {
		cur.Next = &ListNode{Val: nodes[i].Val}
		cur = cur.Next
	}
	return dummy.Next
}

// 迭代
func reverseList3(head *ListNode) *ListNode {
	var pre *ListNode
	// 遍历链表
	for head != nil {
		// 暂存 next 指针
		temp := head.Next
		// 将 Next 指向前一个节点（断开head）
		head.Next = pre
		pre = head
		// 移动
		head = temp
	}
	return pre
}

func main() {
	var head1 *ListNode
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 5)

	TraverseSingleList(head1)
	//rHead := reverseList1(head1)
	//rHead := reverseList2(head1)
	rHead := reverseList3(head1)
	TraverseSingleList(rHead)
}
