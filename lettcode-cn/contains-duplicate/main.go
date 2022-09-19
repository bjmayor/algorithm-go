package main

import "sort"

func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	last := nums[0]
	for _, v := range nums[1:] {
		if last == v {
			return true
		}
		last = v
	}
	return false
}

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
		if m[v] >= 2 {
			return true
		}
	}
	return false
}
