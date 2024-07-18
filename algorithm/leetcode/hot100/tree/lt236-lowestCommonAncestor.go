// [L236-中等] 二叉树的最近公共祖先
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 当越过叶节点，则直接返回 null
	if root == nil {
		return nil
	}
	// 当 root 等于 p,q ，则直接返回 root
	if root == p || root == q {
		return root
	}
	// 分别递归左右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func main() {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)

	node1.Left = node2

	fmt.Println(LevelOrderTraverseTreeNodeToArray(lowestCommonAncestor(node1, node1, node2)))
}
