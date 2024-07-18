// [L543-简单] 二叉树的直径
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

var result int

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left)
	r := depth(root.Right)
	if result < l+r+1 {
		result = l + r + 1
	}
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

func diameterOfBinaryTree(root *TreeNode) int {
	result = 1
	depth(root)
	return result - 1
}

func main() {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)
	node4 := NewTreeNode(4)
	node5 := NewTreeNode(5)

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	fmt.Println(diameterOfBinaryTree(node1))
}
