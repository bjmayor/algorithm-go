package main

import "sort"

func maxProduct(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return (nums[0] - 1) * (nums[1] - 1)
}

func main() {
	println(maxProduct([]int{3, 4, 5, 2}))
}
