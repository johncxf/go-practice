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
	var head1 *ListNode
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 3)
	fmt.Println(printListFromTailToHead(head1))
}
