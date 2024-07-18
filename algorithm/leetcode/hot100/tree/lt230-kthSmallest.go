// [230-中等] 二叉搜索树中第K小的元素
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

// 中序遍历
func kthSmallest1(root *TreeNode, k int) int {
	var list []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		list = append(list, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return list[k-1]
}

// 中序遍历
func kthSmallest2(root *TreeNode, k int) int {
	ans := 0
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if ans != 0 {
			return
		}
		if node == nil {
			return
		}
		inorder(node.Left)
		k--
		if k == 0 {
			ans = node.Val
			return
		}
		inorder(node.Right)
	}
	inorder(root)
	return ans
}

func main() {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)
	node4 := NewTreeNode(4)

	node3.Left = node1
	node3.Right = node4
	node1.Left = node2

	fmt.Println(kthSmallest1(node3, 1))
	fmt.Println(kthSmallest2(node3, 1))
}
