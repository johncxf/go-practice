package datastruct

// DoublyListNode 双向链表
type DoublyListNode struct {
	Val  int             // 节点值
	Next *DoublyListNode // 指向后继节点的指针
	Prev *DoublyListNode // 指向前驱节点的指针
}

// NewDoublyListNode 初始化
func NewDoublyListNode(val int) *DoublyListNode {
	return &DoublyListNode{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}
