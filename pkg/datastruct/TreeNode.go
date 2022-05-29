package datastruct

import (
    "fmt"
)

// 二叉树
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 初始化节点
func NewTreeNode(data int) *TreeNode {
    return &TreeNode{
        Val: data,
        Left: nil,
        Right: nil,
    }
}

// 前序遍历
func PreOrderTraverseTreeNode(treeNode *TreeNode) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}

	// 先打印当前节点值
	fmt.Printf("%v ", treeNode.Val)

	// 然后依次对左子树和右子树做前序遍历
	PreOrderTraverseTreeNode(treeNode.Left)
	PreOrderTraverseTreeNode(treeNode.Right)
}

// 中序遍历
func MidOrderTraverseTreeNode(treeNode *TreeNode) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}

	// 先从左子树最左侧节点开始遍历
	MidOrderTraverseTreeNode(treeNode.Left)
	// 打印位于中间的根节点
	fmt.Printf("%v ", treeNode.Val)
	// 最后按照和左子树一样的逻辑遍历右子树
	MidOrderTraverseTreeNode(treeNode.Right)
}

// 后序遍历
func PostOrderTraverseTreeNode(treeNode *TreeNode) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}

	// 先遍历左子树
	PostOrderTraverseTreeNode(treeNode.Left)
	// 再遍历右子树
	PostOrderTraverseTreeNode(treeNode.Right)
	// 最后访问根节点
	fmt.Printf("%v ", treeNode.Val)
}

// 层次遍历
func LevelOrderTraverseTreeNode(treeNode *TreeNode) {
    if nil == treeNode {
        return
    }
    
    tmp := []*TreeNode{treeNode}
    for i := 0; i < len(tmp); i++ {
        current := tmp[i]

		fmt.Printf("%v ", current.Val)

		if nil != current.Left {
            tmp = append(tmp, current.Left)
        }
        if nil != current.Right {
            tmp = append(tmp, current.Right)
        }
    }
}