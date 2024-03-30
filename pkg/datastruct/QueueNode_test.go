package datastruct

import (
	"testing"
)

func TestTraverseQueue(t *testing.T) {
	q := LinkedQueue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Traverse()
}
