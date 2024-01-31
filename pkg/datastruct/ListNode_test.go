package datastruct

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T) {
	Head = nil
	fmt.Println(AddNode(Head, 1))
	fmt.Println(AddNode(Head, 2))
	fmt.Println(AddNode(Head, 3))
}

func TestTraverseSingleList(t *testing.T) {
	Head = nil
	AddNode(Head, 1)
	AddNode(Head, 2)
	AddNode(Head, 3)

	// 遍历链表
	TraverseSingleList(Head)
}
