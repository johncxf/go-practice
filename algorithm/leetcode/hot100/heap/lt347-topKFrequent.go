// [L347-中等] 前 K 个高频元素
package main

import (
	"container/heap"
	"fmt"
)

type intHeap [][2]int

// Len sort.Interface 的方法，获取堆长度
func (h intHeap) Len() int { return len(h) }

// Less sort.Interface 的方法，定义是小顶堆还是大顶堆
func (h intHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }

// Swap sort.Interface 的方法，交换元素位置
func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push heap.Interface 的方法，实现推入元素到堆
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }

// Pop heap.Interface 的方法，实现弹出堆顶元素
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// hashMap 存储 数字和出现次数对应关系
	hash := make(map[int]int, 0)
	for _, num := range nums {
		hash[num]++
	}
	// 构建小顶堆
	h := &intHeap{}
	heap.Init(h)
	// 遍历 hashMap
	for key, val := range hash {
		// 将当前元素和元素出现次数入堆
		heap.Push(h, [2]int{key, val})
		// 如果堆长度超过k，则移除堆顶
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	// 依次取出堆中元素
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

// 桶排序
func topKFrequent2(nums []int, k int) []int {
	// 先遍历数组统计每个元素出现的次数
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
	}
	// 桶排序
	// 构建 n+1 个桶
	buckets := make([][]int, len(nums)+1)
	// 以元素出现的次数为下标分配到桶里
	for i, v := range hash {
		buckets[v] = append(buckets[v], i)
	}
	// 倒序遍历所有桶，取出 k 个元素
	var ans []int
	for i := len(buckets) - 1; i > 0 && k > 0; i-- {
		if len(buckets[i]) > 0 {
			k = k - len(buckets[i])
			ans = append(ans, buckets[i]...)
		}
	}
	return ans
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

	fmt.Println(topKFrequent2([]int{1, 1, 1, 2, 2, 3}, 2))
}
