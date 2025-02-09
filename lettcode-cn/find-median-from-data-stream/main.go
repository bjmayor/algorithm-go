package main

import (
	"container/heap"
	"fmt"
)

// minHeap 实现最小堆
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// maxHeap 实现最大堆
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
* MedianFinder 数据流的中位数查找器
* 使用两个堆来维护:
* - maxHeap 存储较小的一半数
* - minHeap 存储较大的一半数
* 保持 maxHeap.size >= minHeap.size
 */
type MedianFinder struct {
	minH *minHeap // 最小堆,存储较大的一半数
	maxH *maxHeap // 最大堆,存储较小的一半数
}

/*
* Constructor 初始化 MedianFinder
* @return MedianFinder 初始化的对象
 */
func Constructor() MedianFinder {
	minH := &minHeap{}
	maxH := &maxHeap{}
	heap.Init(minH)
	heap.Init(maxH)
	return MedianFinder{minH: minH, maxH: maxH}
}

/*
* AddNum 添加数字到数据流
* @param num 要添加的数字
* 1. 如果maxHeap为空或num小于maxHeap堆顶,加入maxHeap
* 2. 否则加入minHeap
* 3. 平衡两个堆的大小,保持maxHeap.size >= minHeap.size
 */
func (mf *MedianFinder) AddNum(num int) {
	if mf.maxH.Len() == 0 || num < (*mf.maxH)[0] {
		heap.Push(mf.maxH, num)
	} else {
		heap.Push(mf.minH, num)
	}

	// 平衡两个堆的大小
	if mf.maxH.Len() < mf.minH.Len() {
		heap.Push(mf.maxH, heap.Pop(mf.minH))
	} else if mf.maxH.Len() > mf.minH.Len()+1 {
		heap.Push(mf.minH, heap.Pop(mf.maxH))
	}
}

/*
* FindMedian 查找当前数据流的中位数
* @return float64 中位数
* 如果两个堆大小相等,取两堆顶的平均值
* 否则取maxHeap的堆顶
 */
func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxH.Len() > mf.minH.Len() {
		return float64((*mf.maxH)[0])
	}
	return float64((*mf.maxH)[0]+(*mf.minH)[0]) / 2.0
}

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian()) // 1.5
	obj.AddNum(3)
	fmt.Println(obj.FindMedian()) // 2.0
}
