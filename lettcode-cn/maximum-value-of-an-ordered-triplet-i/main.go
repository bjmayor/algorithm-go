package main

import "fmt"

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], nums[i-1])
	}
	for i := 1; i < n; i++ {
		rightMax[n-1-i] = max(rightMax[n-i], nums[n-i])
	}
	var res int64 = 0
	for j := 1; j < n-1; j++ {
		res = max(res, int64((leftMax[j]-nums[j])*rightMax[j]))
	}
	return res
}

func main() {
	nums := []int{12, 6, 1, 2, 7}
	fmt.Println(maximumTripletValue(nums)) // 77
}
