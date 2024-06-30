package datastruct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTraverseSingleList(t *testing.T) {
	ast := assert.New(t)

	var head *ListNode
	head = AddNode(head, 1)
	head = AddNode(head, 2)
	head = AddNode(head, 3)
	TraverseSingleList(head)
	size := GetSingleListSize(head)
	ast.Equal(size, 3)
}
