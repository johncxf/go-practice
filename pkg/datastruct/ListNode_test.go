package datastruct

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T) {
	head = nil
	fmt.Println(AddNode(head, 1))
	fmt.Println(AddNode(head, 2))
	fmt.Println(AddNode(head, 2))
}

func TestListNode(t *testing.T) {
	head = nil
	AddNode(head, 1)
	AddNode(head, 2)
	AddNode(head, 2)

	// 遍历链表
	TraverseSingleList(head)
}
