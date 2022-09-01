package main

func singleNumber(nums []int) int {
	v := 0
	for _, vv := range nums {
		v ^= vv
	}
	return v
}
