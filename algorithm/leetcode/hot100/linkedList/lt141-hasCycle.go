// [L141-中等] 环形链表
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

// hashMap，遍历链表，将节点存入 hashMap，如果已存在相同节点，则说明存在环
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	hash := map[*ListNode]bool{}
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}
		hash[head] = true
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
