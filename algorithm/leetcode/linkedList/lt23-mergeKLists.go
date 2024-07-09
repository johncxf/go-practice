// L23-困难] 合并 K 个升序链表
package main

import (
	. "go_practice/pkg/datastruct"
)

// 依次合并
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	ans := lists[0]
	for i := 1; i < len(lists); i++ {
		ans = mergeList(ans, lists[i])
	}
	return ans
}

// 分治合并
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge1(lists, 0, len(lists)-1)
}

func merge1(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	return mergeList(merge1(lists, left, mid), merge1(lists, mid+1, right))
}

// 合并两个有序链表
func mergeList(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummy := &ListNode{0, nil}
	cur := dummy
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val < list2.Val {
			cur.Next = list1
			cur = cur.Next
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = cur.Next
			list2 = list2.Next
		}
	}
	return dummy.Next
}

func main() {
	var (
		head1 *ListNode
		head2 *ListNode
		head3 *ListNode
	)
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 5)

	head2 = AddNode(head2, 1)
	head2 = AddNode(head2, 3)
	head2 = AddNode(head2, 4)

	head3 = AddNode(head3, 2)
	head3 = AddNode(head3, 6)

	//TraverseSingleList(mergeKLists1([]*ListNode{head1, head2, head3}))
	TraverseSingleList(mergeKLists2([]*ListNode{head1, head2, head3}))

}
