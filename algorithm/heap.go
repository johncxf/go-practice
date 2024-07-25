// 堆
package main

import (
	"container/heap"
	"fmt"
)

// Go 语言中可以通过实现 heap.Interface 来构建整数大顶堆
// 实现 heap.Interface 需要同时实现 sort.Interface

type intHeap []any

// Push heap.Interface 的方法，实现推入元素到堆
func (h *intHeap) Push(x any) {
	*h = append(*h, x)
}

// Pop heap.Interface 的方法，实现弹出堆顶元素
func (h *intHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// Len sort.Interface 的方法，获取堆长度
func (h intHeap) Len() int { return len(h) }

// Less sort.Interface 的方法，定义是小顶堆还是大顶堆（小顶堆改成 < ）
func (h intHeap) Less(i, j int) bool { return h[i].(int) > h[j].(int) }

// Swap sort.Interface 的方法，交换元素位置
func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Top 获取堆顶元素
func (h *intHeap) Top() any {
	return (*h)[0]
}

func main() {
	/* 初始化堆 */
	// 初始化大顶堆
	maxHeap := &intHeap{}
	heap.Init(maxHeap)
	/* 元素入堆 */
	// 调用 heap.Interface 的方法，来添加元素
	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 4)
	heap.Push(maxHeap, 5)

	/* 获取堆顶元素 */
	top := maxHeap.Top()
	fmt.Printf("堆顶元素为 %d\n", top)

	/* 堆顶元素出堆 */
	// 调用 heap.Interface 的方法，来移除元素
	heap.Pop(maxHeap) // 5
	heap.Pop(maxHeap) // 4
	heap.Pop(maxHeap) // 3
	heap.Pop(maxHeap) // 2
	heap.Pop(maxHeap) // 1

	/* 获取堆大小 */
	size := len(*maxHeap)
	fmt.Printf("堆元素数量为 %d\n", size)

	/* 判断堆是否为空 */
	isEmpty := len(*maxHeap) == 0
	fmt.Printf("堆是否为空 %t\n", isEmpty)
}
