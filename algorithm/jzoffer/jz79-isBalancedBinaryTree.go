package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

/**
 * [JZ79-简单] 判断是否是平衡二叉树
 *
 * @param pRoot TreeNode类
 * @return bool布尔型
 */
func isBalancedBinaryTree(pRoot *TreeNode) bool {
	return treeDepth1(pRoot) > -1
}

func treeDepth1(root *TreeNode) int {
	if nil == root {
		return 0
	}

	leftLength := treeDepth1(root.Left)
	rightLength := treeDepth1(root.Right)

	if -1 == leftLength || -1 == rightLength || 1 < abs(leftLength-rightLength) {
		return -1
	}

	return max(leftLength, rightLength) + 1
}

func abs(a int) int {
	if 0 > a {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)
	node4 := NewTreeNode(4)
	node5 := NewTreeNode(5)
	node6 := NewTreeNode(6)
	node7 := NewTreeNode(7)

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7

	res := isBalancedBinaryTree(node1)
	fmt.Println(res)
}
