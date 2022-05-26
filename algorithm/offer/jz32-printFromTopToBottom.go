package main

import (
	"fmt"
    . "go_practice/pkg/datastruct"
)

/**
 * [JZ32-简单]从上往下打印二叉树
 * 
 * @param root TreeNode类 
 * @return int 整型一维数组
*/
func printFromTopToBottom(root *TreeNode) []int {
    res := []int{}
    if nil == root {
        return res
    }
    
    tmp := []*TreeNode{root}
    for i := 0; i < len(tmp); i++ {
        current := tmp[i]
        res = append(res, current.Val)
        if nil != current.Left {
            tmp = append(tmp, current.Left)
        }
        if nil != current.Right {
            tmp = append(tmp, current.Right)
        }
    }
    return res
}

func main() {
    tree1 := NewTreeNode(0)
    tree2 := NewTreeNode(2)
    tree3 := NewTreeNode(3)

    tree1.Left = tree2
    tree1.Right = tree3

    res := printFromTopToBottom(tree1)
    fmt.Println(res)
}