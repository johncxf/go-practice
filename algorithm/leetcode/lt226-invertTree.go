package main

import (
	. "go_practice/pkg/datastruct"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 交换左右节点位置
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	// 递归
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func main() {
	node1 := NewTreeNode(1) // 根节点
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)
	node4 := NewTreeNode(4)
	node5 := NewTreeNode(5)
	node6 := NewTreeNode(6)
	node7 := NewTreeNode(7)
	node8 := NewTreeNode(8)

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node4.Left = node7
	node4.Right = node8

	rNode := invertTree(node1)
	LevelOrderTraverseTreeNode(rNode)
}
