package main

func minStartValue(nums []int) int {
	sum := 0
	m := sum
	for _, v := range nums {
		sum += v
		if sum < m {
			m = sum
		}
	}
	if m < 0 {
		return 1 - m
	}
	return 1
}
