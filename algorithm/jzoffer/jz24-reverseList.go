package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

/**
 * [JZ24-简单] 反转链表-双链表法
 *
 * @param head ListNode类
 * @return ListNode类
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var newHead *ListNode
	var tmp *ListNode
	for head != nil {
		// 暂存下一个节点
		tmp = head.Next
		// 把新链表挂到原链表后面
		head.Next = newHead
		TraverseSingleList(head)
		// 更新新链表
		newHead = head
		// 重新赋值，继续访问
		head = tmp
	}

	return newHead
}

/**
 * [JZ24-简单] 反转链表-栈
 *
 * @param head ListNode类
 * @return ListNode类
 */
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tmpMap := make([]int, 0)
	for head != nil {
		tmpMap = append([]int{head.Val}, tmpMap...)
		head = head.Next
	}

	rHead := &ListNode{Val: tmpMap[0], Next: nil}
	cHead := rHead
	for i := 1; i < len(tmpMap); i++ {
		tmpHead := &ListNode{Val: tmpMap[i], Next: nil}
		cHead.Next = tmpHead
		cHead = cHead.Next
	}
	return rHead
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
