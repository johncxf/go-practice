// [L114-中等] 二叉树展开为链表
package main

import (
	. "go_practice/pkg/datastruct"
)

// 前序遍历
func flatten(root *TreeNode) {
	// 前序遍历二叉树，存储节点
	list := make([]*TreeNode, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	// 遍历节点列表，修改二叉树左右节点信息
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
	//fmt.Println(LevelOrderTraverseTreeNodeToArray(root))
}

func main() {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)
	node4 := NewTreeNode(4)
	node5 := NewTreeNode(5)
	node6 := NewTreeNode(6)

	node1.Left = node2
	node1.Right = node5
	node2.Left = node3
	node2.Right = node4
	node5.Right = node6

	flatten(node1)
}
