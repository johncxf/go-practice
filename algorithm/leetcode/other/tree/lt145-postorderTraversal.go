package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

// [L145-简单] 二叉树的后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return
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

	fmt.Println(postorderTraversal(node1))
}
