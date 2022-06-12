package datastruct

import (
	"fmt"
	"testing"
)

func TestTreeNode(t *testing.T) {
	node1 := NewTreeNode(1) // 根节点
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)
	node4 := NewTreeNode(4)
	node5 := NewTreeNode(5)
	node6 := NewTreeNode(6)
	node7 := NewTreeNode(7)
	node8 := NewTreeNode(8)

    node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node4.Left = node7
	node4.Right = node8

	// 前序遍历二叉树
	fmt.Print("前序遍历: ")
	PreOrderTraverseTreeNode(node1)
	fmt.Println()

    // 中序遍历二叉树
	fmt.Print("中序遍历: ")
	MidOrderTraverseTreeNode(node1)
	fmt.Println()

	// 后序遍历二叉树
	fmt.Print("后序遍历: ")
	PostOrderTraverseTreeNode(node1)
	fmt.Println()

	// 层次遍历二叉树
	fmt.Print("层次遍历: ")
	LevelOrderTraverseTreeNode(node1)
	fmt.Println()

	// 将数组转化成二叉树
	testNode := TransformArrayToTreeNode([]int{5,4,8,11,NULL,13,4,7,2,NULL,NULL,5,1})
	LevelOrderTraverseTreeNode(testNode)
	fmt.Println()
}