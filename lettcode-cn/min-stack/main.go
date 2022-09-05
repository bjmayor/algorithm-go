package main

import (
	"math"
)

type MinStack struct {
	data    []int
	minData []int
}

func Constructor() MinStack {
	return MinStack{
		minData: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	m := this.minData[len(this.minData)-1]
	this.minData = append(this.minData, min(m, val))
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.minData = this.minData[:len(this.minData)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minData[len(this.minData)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
