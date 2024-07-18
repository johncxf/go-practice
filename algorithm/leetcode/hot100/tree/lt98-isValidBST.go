// [L98-中等] 验证二叉搜索树
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
	"math"
)

func isValidBST(root *TreeNode) bool {
	var inorder func(node *TreeNode)
	// 根据题意，节点最小值为-2的32次方
	tmp := math.MinInt64
	res := true
	inorder = func(node *TreeNode) {
		// 已知结果，直接返回
		if res == false {
			return
		}
		if node == nil {
			return
		}
		inorder(node.Left)
		// 当前节点值必需大于前一个节点值，否则为非二叉搜索树
		if node.Val <= tmp {
			res = false
			return
		}
		tmp = node.Val
		inorder(node.Right)
	}
	inorder(root)
	return res
}

func main() {
	node1 := NewTreeNode(2)
	node2 := NewTreeNode(1)
	node3 := NewTreeNode(3)

	node1.Left = node2
	node1.Right = node3

	fmt.Println(isValidBST(node1))
}
