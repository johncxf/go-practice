package main

import (
	. "go_practice/pkg/datastruct"
)

/**
 * [JZ22-简单] 链表中倒数最后k个结点
 *
 * @param pHead ListNode类
 * @param k int整型
 * @return ListNode类
 */
func findKthToTail(pHead *ListNode, k int) *ListNode {
	res := pHead
	// 获取链表长度
	n := 0
	for pHead != nil {
		n++
		pHead = pHead.Next
	}
	// 链表长度小于 k，则返回空
	if n < k {
		return nil
	}
	for i := 1; i <= (n - k); i++ {
		res = res.Next
	}
	return res
}

func main() {
	var head1 *ListNode
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 5)

	kHead := findKthToTail(head1, 2)
	TraverseSingleList(kHead)
}
