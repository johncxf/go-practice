// 单链表
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

	// if v == t.Value {
	// 	fmt.Println("节点已存在:", v)
	// 	return -1
	// }

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
 * @param head Node类
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
 * 链表中倒数最后 k 个结点
 *
 * @param pHead Node类
 * @param k int整型
 * @return Node类
 */
func findKthToTail(pHead *Node, k int) *Node {
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

/**
 * [JZ18-简单]删除链表的节点
 *
 * @param head Node类
 * @param val int整型
 * @return Node类
 */
func deleteNode(head *Node, val int) *Node {
	if head.Value == val {
		return head.Next
	}
	pre := head
	for nil != head {
		if head.Value == val {
			head.Value = head.Next.Value
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}
	// 简化 - 直接从 head 的 Next 节点开始遍历
	// for head.Next.Value != val {
	// 	head = head.Next
	// }
	// head.Next = head.Next.Next
	return pre
}

/**
 * [JZ24-简单]反转链表
 *
 * @param pHead Node类
 * @return Node类
 */
func reverseList(head *Node) *Node {
	// tmp := make([]int, 0)
	// for nil != head {
	// 	tmp = append(tmp, head.Value)
	// 	head = head.Next
	// }

	// head = new(Node)
	// tmpLen := len(tmp) - 1
	// for tmpLen >= 0 {
	// 	addNode(head, tmp[tmpLen])
	// 	tmpLen--
	// }

	var newHead *Node
	var tmp *Node
	for head != nil {
		// 暂存下一个节点
		tmp = head.Next
		// 把新链表挂到原链表后面
		head.Next = newHead
		// 更新新链表
		newHead = head
		// 移动指针
		head = tmp
	}

	return newHead
}

/**
 * [JZ76-简单]删除链表中重复的结点
 *
 * @param pHead Node类
 * @return Node类
 */
func deleteDuplication(pHead *Node) *Node {
	// 定义一个空链表
	var res = new(Node)
	// 给新链表添加表头
	res.Next = pHead
	cur := res
	// 从 next 节点开始遍历（前面加了一个表头）
	for nil != cur.Next && nil != cur.Next.Next {
		if cur.Next.Value == cur.Next.Next.Value {
			// 遇到前后节点相同情况，进行遍历，跳过所有相同节点
			tmp := cur.Next.Value
			for nil != cur.Next && cur.Next.Value == tmp {
				cur.Next = cur.Next.Next
			}
		} else {
			// 循环链表
			cur = cur.Next
		}
	}
	return res.Next
}

/**
 * [JZ23-中等]链表中环的入口结点
 *
 * @param pHead Node类
 * @return Node类
 */
func entryNodeOfLoop(pHead *Node) *Node {
	if nil == pHead || nil == pHead.Next {
		return nil
	}
	fast, slow := pHead, pHead
	for nil != fast && nil != fast.Next.Next {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = pHead
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}

func main() {
	fmt.Println(head)
	head = nil
	// 遍历链表
	traverse(head)
	// // 添加节点
	// addNode(head, 1)
	// addNode(head, -1)
	// // 再次遍历
	// traverse(head)
	// // 添加更多节点
	// addNode(head, 10)
	// addNode(head, 5)
	// addNode(head, 45)
	// // 添加已存在节点
	// addNode(head, 5)
	// // 再次遍历
	// traverse(head)

	addNode(head, 1)
	addNode(head, 2)
	addNode(head, 2)
	addNode(head, 3)
	addNode(head, 3)
	addNode(head, 5)
	traverse(head)

	// deleteNode(head, 5)

	// ret := reverseList(head)
	// traverse(ret)

	ret := deleteDuplication(head)
	traverse(ret)

	// res := printListFromTailToHead(head)
	// fmt.Println(res)
}
