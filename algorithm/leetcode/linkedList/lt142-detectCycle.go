// [L142-中等] 环形链表 II
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

func detectCycle(head *ListNode) *ListNode {
	hash := map[*ListNode]struct{}{}
	i := 0
	for head != nil {
		if _, ok := hash[head]; ok {
			return head
		}
		i++
		hash[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func main() {
	var head1 *ListNode
	head1 = AddNode(head1, 3)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 0)
	head1 = AddNode(head1, -4)
	fmt.Println(detectCycle(head1))
}
