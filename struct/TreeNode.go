// 二叉树
package struct

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func NewTreeNode(data int) *TreeNode {
    return &TreeNode{
        Val: data,
        Left: nil,
        Right: nil,
    }
}
