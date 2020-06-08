package  main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	q := NewDeQueue(k)
	result := make([]int,0)
	for i,v := range nums {
			if !q.Empty() && i - q.LTop() ==k {
				q.LPop()
			}
		for !q.Empty() {
			if v > nums[q.RTop()] {
				q.RPop()
			} else {
				break
			}
		}
		q.RPush(i)

		if i+1 >= k {
			result = append(result, nums[q.LTop()])
		}
	}
	return result
}

type DeQueue struct {
	data []int //
}

func NewDeQueue(size int) *DeQueue {
	q := DeQueue{}
	q.data = make([]int,size, size)
	return &q
}

func (q *DeQueue)LTop() int  {
	return q.data[0]
}

func (q *DeQueue)RTop() int  {
	return q.data[len(q.data)-1]
}

func (q *DeQueue)LPop() {
	q.data = q.data[1:]
}

func (q *DeQueue)RPop() {
	q.data = q.data[:len(q.data)-1]
}

func (q *DeQueue)Empty() bool  {
	return len(q.data) == 0
}

func (q *DeQueue)RPush(v int)  {
	q.data = ap
}


func arrEqual(a, b []int) bool  {
	if len(a) != len(b) {
		return false
	}
	for i:=0;i<len(a);i++ {
		if a[i]!=b[i] {
			return false
		}
	}
	return true
}

func main() {
	tests := []struct{
		Nums []int
		Target int
		Expected []int
	} {
		{ []int{1,3,-1,-3,5,3,6,7}, 3, []int{3,3,5,5,6,7} },
		{ []int{1}, 1, []int{1} },
		{ []int{1,-1}, 1, []int{1,-1} },
		{ []int{9,10,9,-7,-4,-8,2,-6}, 5, []int{10,10,9,2} },
	}
	for _, t := range tests {
		if !arrEqual(maxSlidingWindow(t.Nums,t.Target), t.Expected) {
			fmt.Println("expected:", t.Expected, " real:", maxSlidingWindow(t.Nums, t.Target))
		}
	}
}