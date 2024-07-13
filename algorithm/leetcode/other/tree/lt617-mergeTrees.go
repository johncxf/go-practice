package main

import (
	. "go_practice/pkg/datastruct"
)

func mergeTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func main() {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(3)
	node3 := NewTreeNode(2)
	node4 := NewTreeNode(5)
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4

	root1 := NewTreeNode(2)
	root2 := NewTreeNode(1)
	root3 := NewTreeNode(3)
	root4 := NewTreeNode(4)
	root5 := NewTreeNode(7)
	root1.Left = root2
	root1.Right = root3
	root2.Right = root4
	root3.Right = root5
	rRoot := mergeTrees(node1, root1)
	LevelOrderTraverseTreeNode(rRoot)
}
