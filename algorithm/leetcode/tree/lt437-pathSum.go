package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func main() {
	node1 := NewTreeNode(10)
	node2 := NewTreeNode(5)
	node3 := NewTreeNode(-3)
	node4 := NewTreeNode(3)
	node5 := NewTreeNode(2)
	node6 := NewTreeNode(11)
	node7 := NewTreeNode(3)
	node8 := NewTreeNode(-2)
	node9 := NewTreeNode(1)

	node1.Left = node2
	node1.Right = node3
	node3.Right = node6
	node2.Left = node4
	node2.Right = node5
	node4.Left = node7
	node4.Right = node8
	node5.Right = node9

	fmt.Println(pathSum(node1, 8))
}
