package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	var size int
	var result = math.MaxInt32
	var sum int
	for i:=0;i<len(nums);i++ {
		sum += nums[i]
		size++
		if sum  >=  s {
			left := sum
			for left - nums[i+1-size] >=s  {
				left -=nums[i+1-size]
				size--
			}
			sum = left
			if size<result {
				result =size
			}
		}
	}
	if result == math.MaxInt32 {
		return 0
	}
	return result
}

func main()  {
	tests := []struct{
		Nums []int
		s int
		Expected int
	} {
		{ []int{2,3,1,2,4,3},7,2},
	}
	for _, t := range tests {
		var real = minSubArrayLen(t.s, t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
