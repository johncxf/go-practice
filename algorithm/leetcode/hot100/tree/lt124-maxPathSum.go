// [L124-困难] 二叉树中的最大路径和
package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
	"math"
)

// 递归
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		left := max1(maxGain(node.Left), 0)
		right := max1(maxGain(node.Right), 0)
		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		path := node.Val + left + right
		// 更新最大值
		maxSum = max1(maxSum, path)
		// 返回节点最大贡献值
		return node.Val + max1(left, right)
	}
	maxGain(root)
	return maxSum
}

func max1(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)

	node1.Left = node2
	node1.Right = node3

	fmt.Println(maxPathSum(node1))
}
