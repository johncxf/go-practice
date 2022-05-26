package datastruct

// 二叉树
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
