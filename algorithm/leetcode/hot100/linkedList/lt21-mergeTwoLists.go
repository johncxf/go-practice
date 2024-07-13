package main

import . "go_practice/pkg/datastruct"

// [L21-简单] 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var res *ListNode
	if list1.Val > list2.Val {
		res = list2
		res.Next = mergeTwoLists(list1, list2.Next)
	} else {
		res = list1
		res.Next = mergeTwoLists(list1.Next, list2)
	}
	return res
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

	TraverseSingleList(mergeTwoLists(head1, head2))
}
