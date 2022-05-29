package main

import (
	"fmt"
    . "go_practice/pkg/datastruct"
)

/**
 * [JZ68-简单]二叉搜索树的最近公共祖先
 *
 * @param root TreeNode类 
 * @param p int整型 
 * @param q int整型 
 * @return int整型
*/
func lowestCommonAncestor(root *TreeNode, p int, q int) int {
    if nil == root {
        return 0
    }
    
    // p q 都小于等于这个节点值，则 p q 都在左子树
    if root.Val > p && root.Val > q {
        return lowestCommonAncestor(root.Left, p, q)
    }

    // p q 都大于等于这个节点值，则 p q 都在右子树
    if root.Val < p && root.Val < q {
        return lowestCommonAncestor(root.Right, p, q)
    }

    return root.Val
}

func main() {
	node1 := NewTreeNode(7)
	node2 := NewTreeNode(1)
	node3 := NewTreeNode(2)
	node4 := NewTreeNode(0)
	node5 := NewTreeNode(4)
	node6 := NewTreeNode(11)
	node7 := NewTreeNode(14)
	node8 := NewTreeNode(3)
	node9 := NewTreeNode(5)

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node5.Left = node8
	node5.Right = node9

	res := lowestCommonAncestor(node1, 1, 12)
	fmt.Println(res)
}