package main

import (
	"fmt"
    . "go_practice/pkg/datastruct"
)

/**
  * [JZ82-简单] 二叉树中和为某一值的路径(一)
  *
  * @param root TreeNode 类 
  * @param sum int 整型 
  * @return bool 布尔型
*/
func hasTreePathSum(root *TreeNode, sum int) bool {
    if nil == root {
		return false
	}
	if nil == root.Left && nil == root.Right && root.Val == sum {
		return true
	}

	return hasTreePathSum(root.Left, sum - root.Val) || hasTreePathSum(root.Right, sum - root.Val)
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

	res := hasTreePathSum(node1, 2)
	fmt.Println(res)
}