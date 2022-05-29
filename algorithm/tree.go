package main

import (
	"fmt"
)

// 通过链表存储二叉树节点信息
type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
}

// 初始化根节点
func NewNode(data interface{}) *TreeNode {
	return &TreeNode{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

// 实现接口方法
func (node *TreeNode) GetData() string {
	return fmt.Sprintf("%v", node.Data)
}

// 前序遍历
func preOrderTraverse(treeNode *TreeNode) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}

	// 否则先打印当前节点值
	fmt.Printf("%s ", treeNode.GetData())

	// 然后依次对左子树和右子树做前序遍历
	preOrderTraverse(treeNode.Left)
	preOrderTraverse(treeNode.Right)
}

// 中序遍历
func midOrderTraverse(treeNode *TreeNode) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}
	// 否则先从左子树最左侧节点开始遍历
	midOrderTraverse(treeNode.Left)
	// 打印位于中间的根节点
	fmt.Printf("%s ", treeNode.GetData())
	// 最后按照和左子树一样的逻辑遍历右子树
	midOrderTraverse(treeNode.Right)
}

// 后序遍历
func postOrderTraverse(treeNode *TreeNode) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}
	// 否则先遍历左子树
	postOrderTraverse(treeNode.Left)
	// 再遍历右子树
	postOrderTraverse(treeNode.Right)
	// 最后访问根节点
	fmt.Printf("%s ", treeNode.GetData())
}

/**
 * 二叉树的镜像
 *
 * @param pRoot TreeNode 类 
 * @return TreeNode 类
*/
func mirror(pRoot *TreeNode) *TreeNode {
    if nil == pRoot {
        return pRoot
    }
    pRoot.Left, pRoot.Right = mirror(pRoot.Right), mirror(pRoot.Left)
    return pRoot
}

// 测试代码
func main() {
	// 初始化一个简单的二叉树
	node1 := NewNode(0) // 根节点
	node2 := NewNode("1")
	node3 := NewNode(2.0)
	node1.Left = node2
	node1.Right = node3

	// 前序遍历这个二叉树
	fmt.Print("前序遍历: ")
	preOrderTraverse(node1)
	fmt.Println()

	// // 中序遍历这个二叉树
	// fmt.Print("中序遍历: ")
	// midOrderTraverse(node1)
	// fmt.Println()

	// // 后序遍历这个二叉树
	// fmt.Print("后序遍历: ")
	// postOrderTraverse(node1)
	// fmt.Println()


	// mirror(node1)
	// preOrderTraverse(node1)
}