// [L437-中等] 路径总和 III
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := 0
	var preorder func(node *TreeNode, res *int)
	preorder = func(node *TreeNode, res *int) {
		if node == nil {
			return
		}
		*res += rootSum(node, targetSum)
		preorder(node.Left, res)
		preorder(node.Right, res)
	}
	// 遍历所有节点，以每个节点为根节点统计对应符合条件的数量
	preorder(root, &ans)
	return ans
}

// 先序遍历：以当前节点为根节点下所有节点和为 target 的数量
func rootSum(node *TreeNode, target int) (res int) {
	if node == nil {
		return
	}
	v := node.Val
	if v == target {
		res++
	}
	res += rootSum(node.Left, target-v)
	res += rootSum(node.Right, target-v)
	return
}

func main() {
	node1 := NewTreeNode(10)
	node2 := NewTreeNode(5)
	node3 := NewTreeNode(-3)
	node4 := NewTreeNode(3)
	node5 := NewTreeNode(2)
	node6 := NewTreeNode(11)
	node7 := NewTreeNode(3)
	node8 := NewTreeNode(-2)
	node9 := NewTreeNode(1)

	node1.Left = node2
	node1.Right = node3
	node3.Right = node6
	node2.Left = node4
	node2.Right = node5
	node4.Left = node7
	node4.Right = node8
	node5.Right = node9

	fmt.Println(pathSum(node1, 8))
}
