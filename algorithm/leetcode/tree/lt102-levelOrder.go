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
		ret = append(ret, level)
		// 更新队列为下一层的节点
		q = p
	}
	return ret
}

func main() {
	node1 := NewTreeNode(3)
	node2 := NewTreeNode(9)
	node3 := NewTreeNode(20)
	node4 := NewTreeNode(15)
	node5 := NewTreeNode(7)

	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5

	fmt.Println(levelOrder(node1))
}
