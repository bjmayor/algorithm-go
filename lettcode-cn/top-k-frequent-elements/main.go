package main

import (
	"container/heap"
	"fmt"
)

/*
* topKFrequent 计算数组中出现频率最高的 k 个元素
* 算法思想:
* 1. 使用哈希表记录每个元素出现的次数
* 2. 使用堆(优先队列)维护出现次数最多的前 k 个元素
* 3. 返回堆中的元素
 */

// Item 存储元素和其出现次数
type Item struct {
	num   int
	count int
}

// MaxHeap 实现最大堆
type MaxHeap []Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// 使用哈希表记录每个元素出现的次数
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	
	// 使用最大堆维护出现次数最多的前 k 个元素
	h := &MaxHeap{}
	heap.Init(h)
	for num, count := range countMap {
		heap.Push(h, Item{num: num, count: count})
	}
	
	// 返回堆中的前k个元素
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(Item).num
	}
	return result
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
