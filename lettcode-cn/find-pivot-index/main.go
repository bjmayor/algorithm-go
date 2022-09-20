package main

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	left := 0
	right := sum
	for i, v := range nums {
		if i > 0 {
			left += nums[i-1]
		}
		right -= v
		if left == right {
			return i
		}
	}
	return -1
}
