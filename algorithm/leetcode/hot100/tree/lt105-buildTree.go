// [L105-中等] 从前序与中序遍历序列构造二叉树
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

// 递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 已知前序遍历第一个元素为根节点
	root := &TreeNode{Val: preorder[0]}
	// 在中序遍历列表中找到根节点位置
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 左子树的长度
	l := len(inorder[:i])
	// 递归遍历
	root.Left = buildTree(preorder[1:l+1], inorder[:i])
	root.Right = buildTree(preorder[l+1:], inorder[i+1:])
	return root
}

func main() {
	root1 := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(LevelOrderTraverseTreeNodeToArray(root1))
}
