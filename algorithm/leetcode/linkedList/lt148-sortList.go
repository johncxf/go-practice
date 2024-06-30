// [L148-中等] 排序链表
package main

import . "go_practice/pkg/datastruct"

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 定义快慢指针
	slow, fast := head, head.Next
	// 慢指针每次走1步，快指针每次走2步，当快指针fast走到最后的时候，慢指针的位置就是中点
	// 当链表为偶数是，slow指向中间前一个元素
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 暂存
	tmp := slow.Next
	// 切断链表
	slow.Next = nil
	// 递归左右链表，直到切分每个链表剩一个节点终止
	left := sortList(head)
	right := sortList(tmp)

	// 返回合并的结果
	return merge(left, right)
}

// 合并左右链表
func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{} // 创建一个哑节点作为新链表的头
	curr := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			curr.Next = left
			left = left.Next
		} else {
			curr.Next = right
			right = right.Next
		}
		curr = curr.Next
	}

	// 链接剩余的部分
	if left != nil {
		curr.Next = left
	}
	if right != nil {
		curr.Next = right
	}

	// 返回新链表的头（跳过哑节点）
	return dummy.Next
}

func main() {
	var head1 *ListNode
	head1 = AddNode(head1, 4)
	head1 = AddNode(head1, 2)
	head1 = AddNode(head1, 1)
	head1 = AddNode(head1, 3)

	TraverseSingleList(sortList(head1))
}
