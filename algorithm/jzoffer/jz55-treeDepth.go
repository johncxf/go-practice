package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

/**
 * [JZ55-简单] 二叉树的深度（层次遍历）
 *
 * @param pRoot TreeNode类
 * @return int整型
 */
func treeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	tmp := []*TreeNode{pRoot}
	depth := 0
	for i := 0; i < len(tmp); i++ {
		current := tmp[i]
		if nil != current.Left {
			tmp = append(tmp, current.Left)
		}
		if nil != current.Right {
			depth++
			tmp = append(tmp, current.Right)
		}
	}
	return depth
}

func main() {
	node1 := NewTreeNode(1) // 根节点
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
	node5.Left = node7
	depth := treeDepth(node1)
	fmt.Println(depth)
}
