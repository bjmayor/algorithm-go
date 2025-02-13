package main

import (
	"fmt"
	"math/rand/v2"
)

type RandomizedSet struct {
	nums []int
	idx  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: []int{},
		idx:  make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idx[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.idx[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.idx[val]; !ok {
		return false
	}
	idx := this.idx[val]
	last := this.nums[len(this.nums)-1]
	
	// 如果要删除的元素就是最后一个元素,直接删除即可
	if idx == len(this.nums)-1 {
		this.nums = this.nums[:len(this.nums)-1]
		delete(this.idx, val)
		return true
	}
	
	// 如果不是最后一个元素,则需要交换
	this.nums[idx] = last
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.idx, val)
	this.idx[last] = idx
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.IntN(len(this.nums))]
}

func main() {
	randomizedSet := Constructor()
	fmt.Println(randomizedSet.Remove(0))
	fmt.Println(randomizedSet.Remove(0))
	fmt.Println(randomizedSet.Insert(0))
	fmt.Println(randomizedSet.GetRandom())
	fmt.Println(randomizedSet.Remove(0))
	fmt.Println(randomizedSet.Insert(0))
}
