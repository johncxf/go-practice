// [L138-中等] 随机链表的复制
package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 创建一个哈希表，key是原节点，value是新节点
	hash := make(map[*Node]*Node, 0)
	p := head
	// 将原节点和新节点放入哈希表中
	for p != nil {
		newNode := &Node{Val: p.Val}
		hash[p] = newNode
		p = p.Next
	}
	p = head
	// 遍历原链表，设置新节点的next和random
	for p != nil {
		// 获取新节点
		newNode := hash[p]
		// 原节点的下一个节点存在，则更新新节点的下一个节点
		if p.Next != nil {
			newNode.Next = hash[p.Next]
		}
		// 原节点的随机节点存在，则更新新节点的随机节点
		if p.Random != nil {
			newNode.Random = hash[p.Random]
		}
		p = p.Next
	}
	// 返回头节点，即原节点对应的 value 值
	return hash[head]
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}

	node1.Next = node2
	node1.Random = node2
	node2.Next = nil
	node2.Random = node2

	res := copyRandomList(node1)
	for res != nil {
		fmt.Println(res.Val, res.Random)
		res = res.Next
	}
}
