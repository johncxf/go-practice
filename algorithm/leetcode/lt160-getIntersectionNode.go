package main

import (
	. "go_practice/pkg/datastruct"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func main() {
	var head1, head2 *ListNode
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 8)
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 5)

	head2 = AddNode(head2, 5)
	head2 = AddNode(head2, 6)
	head2 = AddNode(head2, 1)
	head2 = AddNode(head2, 8)
	head2 = AddNode(head2, 4)
	head2 = AddNode(head2, 5)

	headK := getIntersectionNode(head1, head2)
	TraverseSingleList(headK)
}
