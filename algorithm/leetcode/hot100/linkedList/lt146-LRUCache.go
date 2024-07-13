// [L146-中等] LRU 缓存
package main

import "fmt"

// DLinkedNode 定义双向链表结构体
type DLinkedNode struct {
	key, value int          // 键值对
	prev, next *DLinkedNode // 指向前驱、后继节点的指针
}

// newDLinkedNode 初始化链表
func newDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

// LRUCache 缓存结构
type LRUCache struct {
	size       int                  // 当前已用容量大小
	capacity   int                  // 容量
	cache      map[int]*DLinkedNode // 哈希表，存储缓存数据的键映射到双向链表的位置
	head, tail *DLinkedNode         // 头部和尾部节点
}

// Constructor 构造函数
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     newDLinkedNode(0, 0),
		tail:     newDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// Get 获取缓存值
func (t *LRUCache) Get(key int) int {
	// 不存在，则返回-1
	if _, ok := t.cache[key]; !ok {
		return -1
	}
	// 获取该值在双向链表中的位置
	node := t.cache[key]
	// 先删除该节点，再添加至链表头部（使访问过的节点处于最新状态）
	t.moveToHead(node)
	// 返回节点的值
	return node.value
}

// Put 插入或更新缓存值
func (t *LRUCache) Put(key int, value int) {
	if node, ok := t.cache[key]; ok { // 已存在，则更新
		node.value = value
		// 先删除该节点，再添加至链表头部（使访问过的节点处于最新状态）
		t.moveToHead(node)
	} else { // 不存在
		// 以该 key value 初始化一个节点
		node = newDLinkedNode(key, value)
		// 加入hash表映射
		t.cache[key] = node
		// 添加该节点至头部
		t.addToHead(node)
		t.size++
		// 若加入该缓存后容量超过最大限制，则删除最后一个缓存
		if t.size > t.capacity {
			rmNode := t.removeTail()    // 删除链表最后一个节点
			delete(t.cache, rmNode.key) // 删除该节点在hash表中的映射关系
			t.size--
		}
	}
}

// 添加节点至头部
func (t *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = t.head
	node.next = t.head.next
	t.head.next.prev = node
	t.head.next = node
}

// 删除节点
func (t *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 删除节点，并将该节点添加至头部
func (t *LRUCache) moveToHead(node *DLinkedNode) {
	t.removeNode(node)
	t.addToHead(node)
}

// 删除链表最后一个节点
func (t *LRUCache) removeTail() *DLinkedNode {
	node := t.tail.prev
	t.removeNode(node)
	return node
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}
