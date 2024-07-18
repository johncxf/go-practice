// [L19-中等] 删除链表的倒数第N个节点
package main

import (
	. "go_practice/pkg/datastruct"
)

// 先遍历链表，获取链表长度，计算出删除第 k 个节点，再次遍历，删除第 k 个节点
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// 先遍历链表，获取链表长度
	length := 0
	tmp := head
	for tmp != nil {
		length++
		tmp = tmp.Next
	}
	// 计算删除第 k 个节点
	k := length - n
	// 如果是头节点，则直接返回后面的节点
	if k == 0 {
		return head.Next
	}
	// 对第 k 个节点进行删除操作
	i := 1
	cur := head
	for cur != nil {
		if i == k {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
		i++
	}
	return head
}

// 快慢指针
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := head, dummy
	// fast 指针先走 n 步，完成后 fast 和 slow 相差 n 个节点
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// fast 和 slow 同时移动
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 此时，删除下一个节点
	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {
	var (
		head1 *ListNode
		head2 *ListNode
	)
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 5)

	head2 = AddNode(head2, 1)

	//TraverseSingleList(removeNthFromEnd1(head1, 2))
	//TraverseSingleList(removeNthFromEnd1(head2, 1))

	TraverseSingleList(removeNthFromEnd2(head1, 2))
	TraverseSingleList(removeNthFromEnd2(head2, 1))
}
