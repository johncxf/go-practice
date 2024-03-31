package datastruct

import (
	"testing"
)

func TestTraverseSingleList(t *testing.T) {
	var head *ListNode
	head = AddNode(head, 1)
	head = AddNode(head, 2)
	head = AddNode(head, 3)
	TraverseSingleList(head)
}
