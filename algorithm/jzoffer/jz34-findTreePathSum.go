package main

import (
	"fmt"
    . "go_practice/pkg/datastruct"
)

var ret [][]int

func dfs(root *TreeNode, expectNumber int, path []int) {
	if nil == root {
		return
	}
	if nil == root.Left && nil == root.Right && root.Val == expectNumber {
		ret = append(ret, append(path, root.Val))
		return
	}
	dfs(root.Left, expectNumber - root.Val, append(path, root.Val))
	dfs(root.Right, expectNumber - root.Val, append(path, root.Val))
}

/**
  * [JZ34-简单] 二叉树中和为某一值的路径(二)
  *
  * @param root TreeNode 类 
  * @param sum int 整型 
  * @return bool 布尔型
*/
func findTreePathSum(root *TreeNode, expectNumber int) [][]int {
	dfs(root, expectNumber, []int{})
	return ret
}

func main() {
	// node1 := NewTreeNode(10)
	// node2 := NewTreeNode(5)
	// node3 := NewTreeNode(12)
	// node4 := NewTreeNode(4)
	// node5 := NewTreeNode(7)

	// node1.Left = node2
	// node1.Right = node3
	// node2.Left = node4
	// node2.Right = node5
	
	arr := []int{5,4,8,11,NULL,13,4,7,2,NULL,NULL,5,1}
	node1 := TransformArrayToTreeNode(arr)

	res := findTreePathSum(node1, 22)
	fmt.Println(res)
}