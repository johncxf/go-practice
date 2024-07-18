// [L199-中等] 二叉树的右视图
package main

import (
	. "go_practice/pkg/datastruct"
)

// 层次遍历，每次取最后一个元素
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	// 层次遍历
	q := []*TreeNode{root}
	for len(q) > 0 {
		// 收集当前层的所有值
		level := make([]int, 0)
		p := make([]*TreeNode, 0)
		for _, node := range q {
			level = append(level, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		// 将收集到的值添加到结果集中
		ans = append(ans, level[len(level)-1])
		// 更新队列为下一层的节点
		q = p
	}
	return ans
}

func main() {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)
	node4 := NewTreeNode(4)
	node5 := NewTreeNode(5)

	node1.Left = node2
	node1.Right = node3
	node2.Right = node5
	node3.Right = node4

	rightSideView(node1)
}
