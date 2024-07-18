// [L21-简单] 合并两个有序链表
package main

import . "go_practice/pkg/datastruct"

// 递归
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var res *ListNode
	if list1.Val > list2.Val {
		res = list2
		res.Next = mergeTwoLists1(list1, list2.Next)
	} else {
		res = list1
		res.Next = mergeTwoLists1(list1.Next, list2)
	}
	return res
}

// 迭代
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	current := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			current = current.Next
			list1 = list1.Next
		} else {
			current.Next = list2
			current = current.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}
	return dummy.Next
}

func main() {
	var (
		head1 *ListNode
		head2 *ListNode
	)
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 4)

	head2 = AddNode(head2, 1)
	head2 = AddNode(head2, 3)
	head2 = AddNode(head2, 4)

	//TraverseSingleList(mergeTwoLists1(head1, head2))
	TraverseSingleList(mergeTwoLists2(head1, head2))
}
