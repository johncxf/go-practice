package main

import (
	"fmt"
)

// 定义节点
type Node struct {
	Value int
	Next  *Node
}

// 初始化头结点
var head = new(Node)

// 添加节点
func addNode(t *Node, v int) int {
	if head == nil {
		t = &Node{v, nil}
		head = t
		return 0
	}

	if v == t.Value {
		fmt.Println("节点已存在:", v)
		return -1
	}

	// 如果当前节点下一个节点为空
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	// 如果当前节点下一个节点不为空
	return addNode(t.Next, v)
}

// 遍历链表
func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> 空链表!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

// 查找节点
func searchNode(t *Node, v int) bool {
	if head == nil {
		t = &Node{v, nil}
		head = t
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return searchNode(t.Next, v)
}

// 获取链表长度
func size(t *Node) int {
	if t == nil {
		fmt.Println("-> 空链表!")
		return 0
	}
	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

/**
 * 从尾到头打印链表
 *
 * @param head ListNode类
 * @return int整型一维数组
 */
func printListFromTailToHead(head *Node) []int {
	ret := make([]int, 0)

	// 遍历链表
	tmp := make([]int, 0)
	for head != nil {
		tmp = append(tmp, head.Value)
		head = head.Next
	}

	// 遍历数组，倒序处理
	tmpLen := len(tmp) - 1
	for tmpLen >= 0 {
		ret = append(ret, tmp[tmpLen])
		tmpLen--
	}

	return ret
}

/**
 * 链表中倒数最后k个结点
 *
 * @param pHead ListNode类
 * @param k int整型
 * @return ListNode类
 */
func FindKthToTail(pHead *Node, k int) *Node {
	// 边界值
	if pHead == nil || k <= 0 {
		return nil
	}

	current, prev := pHead, pHead

	// current 指针先走 k 步
	for i := 0; i < k; i++ {
		if current == nil {
			return nil
		}
		current = current.Next
	}

	// 双指针每轮都向前走一步，直至 current 走过链表 尾节点 时跳出
	// 跳出后，prev 与尾节点距离为 k-1，即 prev 指向倒数第 k 个节点
	for current != nil {
		current = current.Next
		prev = prev.Next
	}
	return prev
}

func main() {
	fmt.Println(head)
	head = nil
	// 遍历链表
	traverse(head)
	// 添加节点
	addNode(head, 1)
	addNode(head, -1)
	// 再次遍历
	traverse(head)
	// 添加更多节点
	addNode(head, 10)
	addNode(head, 5)
	addNode(head, 45)
	// 添加已存在节点
	addNode(head, 5)
	// 再次遍历
	traverse(head)

	res := printListFromTailToHead(head)
	fmt.Println(res)
}
