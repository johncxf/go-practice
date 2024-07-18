// [L104-简单] 二叉树的最大深度
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	node1 := NewTreeNode(1) // 根节点
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)
	node4 := NewTreeNode(4)

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = NewTreeNode(5)
	node3.Left = NewTreeNode(6)
	node4.Left = NewTreeNode(7)
	node4.Right = NewTreeNode(8)

	fmt.Println(maxDepth(node1))
}
