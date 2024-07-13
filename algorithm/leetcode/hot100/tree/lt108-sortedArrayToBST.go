// [L108-简单] 将有序数组转化为二叉搜索树
package main

import (
	. "go_practice/pkg/datastruct"
)

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}

func main() {
	node := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	LevelOrderTraverseTreeNode(node)
}
