// [L25-困难] K 个一组翻转链表
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 哑节点
	dummy := &ListNode{0, head}
	// 初始化 pre 指针
	pre := dummy
	for pre != nil {
		// 尾部指针
		tail := pre
		// 尾部指针 tail 每次移动 k 步
		for i := 0; i < k; i++ {
			tail = tail.Next
			// 剩余可移动步数小于k，则直接返回
			if tail == nil {
				return dummy.Next
			}
		}
		// 暂存尾部指针后面的节点，然后将指针断开
		nex := tail.Next
		tail.Next = nil
		// 定义 start 指针，作为待反转链表的头部
		start := pre.Next
		// 进行反转链表，反转后重新连接到 pre 上
		pre.Next = reverse(start)
		// 将前面断开的部分重新连接上（此时的 start 是该段链表的尾部）
		start.Next = nex
		// pre 指针移动到当前已反转完成的尾部
		pre = start
	}
	return dummy.Next
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
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
	fmt.Println("反转后：")
	TraverseSingleList(reverseKGroup(head1, 2))
}
