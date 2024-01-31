package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

/**
 * [JZ76-中等] 删除链表中重复的结点
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func deleteDuplication(pHead *ListNode) *ListNode {
	// 定义一个空链表
	var res = new(ListNode)
	// 给新链表添加表头
	res.Next = pHead
	cur := res
	// 从 next 节点开始遍历（前面加了一个表头）
	for nil != cur.Next && nil != cur.Next.Next {
		if cur.Next.Val == cur.Next.Next.Val {
			// 遇到前后节点相同情况，进行遍历，跳过所有相同节点
			tmp := cur.Next.Val
			for nil != cur.Next && cur.Next.Val == tmp {
				cur.Next = cur.Next.Next
			}
		} else {
			// 循环链表
			cur = cur.Next
		}
	}
	return res.Next
}

func main() {
	fmt.Println("")
	var head1 *ListNode

	head1 = AddNode2(head1, 1)
	head1 = AddNode2(head1, 2)
	head1 = AddNode2(head1, 3)
	head1 = AddNode2(head1, 3)
	head1 = AddNode2(head1, 4)
	head1 = AddNode2(head1, 4)
	head1 = AddNode2(head1, 5)

	rHead := deleteDuplication(head1)
	TraverseSingleList(rHead)
}
