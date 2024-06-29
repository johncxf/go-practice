package datastruct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	ast := assert.New(t)

	q := NewQueue()
	//fmt.Println(q.IsEmpty())
	ast.True(q.IsEmpty())

	q.Push(1)
	q.Push(2)
	q.Push(3)
	//fmt.Println(q.Len())
	ast.Equal(3, q.Len())

	q.Pop()
	//fmt.Println(q.Len())
	ast.Equal(q.Len(), 2)
}
