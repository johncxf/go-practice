package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	hash := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}
		hash[head] = struct{}{}
		head = head.Next
	}
	return false
}

func main() {
	var head1 *ListNode
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 0)
	head1 = AddNode(head1, -4)
	fmt.Println(hasCycle(head1))
}
