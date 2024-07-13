// [L295-困难] 数据流的中位数
package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *IntHeap) Pop() interface{} {
	tmp := *h
	n := len(tmp)
	x := tmp[n-1]
	*h = tmp[:n-1]
	return x
}

type MedianFinder struct {
	// queMin：存储小于中位数的值（大顶堆），queMax：存储大于中位数的值（小顶堆）
	// queMin 比 queMax 多存一个元素，或者一样多
	queMin, queMax IntHeap
}

func Constructor() MedianFinder {
	h := MedianFinder{}
	heap.Init(&h.queMin)
	heap.Init(&h.queMax)
	return h
}

func (mf *MedianFinder) AddNum(num int) {
	// num 小于等于中位数时，添加数据到 queMin 中
	// 这里使用-的原因时，这里只定义了小顶堆结构，所以使用负数入队表示大顶堆
	if mf.queMin.Len() == 0 || num <= -mf.queMin[0] {
		heap.Push(&mf.queMin, -num)
		// 当 mf.queMin 元素比 mf.queMax 多2个，则添加到 mf.queMax 中
		if mf.queMax.Len()+1 < mf.queMin.Len() {
			heap.Push(&mf.queMax, -heap.Pop(&mf.queMin).(int))
		}
	} else {
		heap.Push(&mf.queMax, num)
		// 如果 mf.queMax 元素比 mf.queMin 多
		if mf.queMax.Len() > mf.queMin.Len() {
			heap.Push(&mf.queMin, -heap.Pop(&mf.queMax).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.queMin.Len() > mf.queMax.Len() {
		return float64(-mf.queMin[0])
	}
	return float64(mf.queMax[0]-mf.queMin[0]) / 2
}

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
}
