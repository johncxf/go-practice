// [L25-困难] K 个一组翻转链表
package main

import (
	. "go_practice/pkg/datastruct"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	res := &ListNode{Next: head}
	pre := res
	for head != nil {
		tail := pre
		// 移动 k 步
		for i := 0; i < k; i++ {
			tail = tail.Next
			// 这一轮剩余步数小于k，直接返回
			if tail == nil {
				return res.Next
			}
		}
		// 暂存节点
		temp := tail.Next
		// 反转
		head, tail = reverse(head, tail)
		// 重新连接首位节点
		pre.Next = head
		tail.Next = temp
		// 下一轮
		pre = tail
		head = tail.Next
	}
	return res.Next
}

// 反转链表
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		temp := p.Next
		p.Next = prev
		prev = p
		p = temp
	}
	return tail, head
}

func main() {
	var head1 *ListNode
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 5)

	TraverseSingleList(reverseKGroup(head1, 2))
}
