package main

import "sort"

func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	sum := 0
	for _, v := range nums {
		sum +=v
	}
	for i,s:=0,0;i<len(nums);i++ {
		s += nums[i]
		if s > sum-s {
			return nums[:i+1]
		}
	}
	return nil
}
