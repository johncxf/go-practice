package datastruct

import (
	"fmt"
	"testing"
)

func TestDoublyListNode(t *testing.T) {
	head := NewDoublyListNode(1)
	fmt.Println(head.Val)
}
