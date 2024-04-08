package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

// [L101-简单] 对称二叉树
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

func main() {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(2)

	node1.Left = node2
	node1.Right = node3
	node2.Left = NewTreeNode(3)
	node2.Right = NewTreeNode(4)
	node3.Left = NewTreeNode(4)
	node3.Right = NewTreeNode(3)

	fmt.Println(isSymmetric(node1))
}
