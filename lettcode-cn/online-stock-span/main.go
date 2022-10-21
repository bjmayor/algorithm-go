package main

import "math"

type StockSpanner struct {
	data [][2]int
	idx  int
}

func Constructor() StockSpanner {
	return StockSpanner{[][2]int{{-1, math.MaxInt32}}, 0}
}

func (this *StockSpanner) Next(price int) int {
	for len(this.data) > 0 {
		if this.data[len(this.data)-1][1] <= price {
			this.data = this.data[:len(this.data)-1]
		} else {
			break
		}
	}
	top := this.data[len(this.data)-1][0]
	this.data = append(this.data, [2]int{this.idx, price})
	ans := this.idx - top
	this.idx++
	return ans
}
