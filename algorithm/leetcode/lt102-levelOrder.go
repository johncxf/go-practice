package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

// [L102-简单] 二叉树的层次遍历
func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := make([]*TreeNode, 0)
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func main() {
	//node1 := NewTreeNode(1) // 根节点
	//node2 := NewTreeNode(2)
	//node3 := NewTreeNode(3)
	//node4 := NewTreeNode(4)
	//node5 := NewTreeNode(5)
	//node6 := NewTreeNode(6)
	//node7 := NewTreeNode(7)
	//node8 := NewTreeNode(8)
	//
	//node1.Left = node2
	//node1.Right = node3
	//node2.Left = node4
	//node2.Right = node5
	//node3.Left = node6
	//node4.Left = node7
	//node4.Right = node8

	node1 := NewTreeNode(3)
	node2 := NewTreeNode(9)
	node3 := NewTreeNode(2)
	node4 := NewTreeNode(1)
	node5 := NewTreeNode(7)

	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5
	fmt.Println(levelOrder(node1))
	//levelOrder(node1)
}
