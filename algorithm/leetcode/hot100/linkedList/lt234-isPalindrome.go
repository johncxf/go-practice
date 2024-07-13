package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	tmp := make([]int, 0)
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	n := len(tmp)
	for i := 0; i < n/2; i++ {
		if tmp[i] != tmp[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var head1, head2 *ListNode
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 1)

	head2 = AddNode(head2, 1)
	head2 = AddNode(head2, 2)

	fmt.Println(isPalindrome(head1))
	fmt.Println(isPalindrome(head2))
}
