// [L104-简单] 二叉树的最大深度
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

// 深度优先搜索
// - 时间复杂度：O(N)
// - 空间复杂度：O(height) 二叉树深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 广度优先搜索
// - 时间复杂度：O(N)
// - 空间复杂度：O(N)
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		p := make([]*TreeNode, 0)
		for _, node := range q {
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		ans++
		q = p
	}
	return ans
}

func main() {
	node1 := NewTreeNode(1) // 根节点
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)
	node4 := NewTreeNode(4)

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = NewTreeNode(5)
	node3.Left = NewTreeNode(6)
	node4.Left = NewTreeNode(7)
	node4.Right = NewTreeNode(8)

	fmt.Println(maxDepth(node1))
	fmt.Println(maxDepth2(node1))
}
