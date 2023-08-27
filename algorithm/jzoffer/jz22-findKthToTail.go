package main

import (
	"fmt"
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
	fmt.Println("")
	Head = nil
	AddNode(Head, 1)
	AddNode(Head, 2)
	AddNode(Head, 3)
	AddNode(Head, 4)
	AddNode(Head, 5)

	kHead := findKthToTail(Head, 2)
	TraverseSingleList(kHead)
}
