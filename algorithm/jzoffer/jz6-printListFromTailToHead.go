package main

import (
	"fmt"
	. "go_practice/pkg/datastruct"
)

/**
 * [JZ6-简单] 从尾到头打印链表
 *
 * @param head ListNode类
 * @return int整型一维数组
 */
func printListFromTailToHead(head *ListNode) []int {
	// 定义切片
	ret := make([]int, 0)
	// 遍历链表
	for head != nil {
		// 将链表值插入切片数组头部
		ret = append([]int{head.Val}, ret...)
		head = head.Next
	}

	return ret
}

func main() {
	fmt.Println("")
	Head = nil
	AddNode(Head, 1)
	AddNode(Head, 2)
	AddNode(Head, 3)
	fmt.Println(printListFromTailToHead(Head))
}
