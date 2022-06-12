package main

import (
	"fmt"
    . "go_practice/pkg/datastruct"
)

func recursion(root1 *TreeNode, root2 *TreeNode) bool {
	if nil == root1 && nil == root2 {
		return true
	}
	if nil == root1 || nil == root2 || root1.Val != root2.Val {
		return false
	}
	return recursion(root1.Left, root2.Right) && recursion(root1.Right, root2.Left)
}

/**
 * [JZ28-简单] 对称二叉树
 *
 * @param pRoot TreeNode 类
 * @return bool 布尔型
*/
func isSymmetricalTree(pRoot *TreeNode) bool {
    return recursion(pRoot, pRoot)
}

func main() {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)
	// node3 := NewTreeNode(3)
	// node4 := NewTreeNode(4)
	// node5 := NewTreeNode(5)
	// node6 := NewTreeNode(6)
	// node7 := NewTreeNode(7)

	node1.Left = node2
	node1.Right = node2
	// node2.Left = node4
	// node2.Right = node5
	// node3.Left = node6
	// node3.Right = node7

	res := isSymmetricalTree(node1)
	fmt.Println(res)
}