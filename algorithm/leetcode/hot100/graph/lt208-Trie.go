// [L208-中等] 实现 Trie (前缀树)
package main

import "fmt"

// Trie 前缀树结构体
type Trie struct {
	// 指向子节点的指针数组
	// children[0]对应小写字母a，children[1]对应小写字母b，…，children[25]对应小写字母z
	children [26]*Trie
	// 该节点是否为字符串结尾
	isEnd bool
}

// Constructor 构造函数
func Constructor() Trie {
	return Trie{}
}

// Insert 插入单词
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		// 减去小写字母a的ASCII值（97）：a->0，b->1，...，z->25
		ch -= 'a'
		// 子节点不存在，则创建一个新的节点
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		// 沿着指针移动到子节点，继续处理下一个字符
		node = node.children[ch]
	}
	// 处理到最后一个字符，标记为结束位置
	node.isEnd = true
}

// Search 查找单词
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

// StartsWith 单词是否是这个前缀
func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

// SearchPrefix 查找前缀
func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	// 遍历前缀所有单词，依次查找
	for _, ch := range prefix {
		ch -= 'a'
		// 有一个单词不存在，则直接返回 nil
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func main() {
	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))
	obj.Insert("app")
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.Search("apple"))
}
