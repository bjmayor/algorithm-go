package main

import "sort"

func containsDuplicate(nums []int) bool {
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
