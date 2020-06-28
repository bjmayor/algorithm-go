package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left :=0
	right := len(nums)-1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid+1
		} else {
			right = mid-1
		}
	}
	return left
}

func main()  {
	tests := []struct{
		Nums []int
		Target int
		Expected int
	} {
		{ []int{1,3,5,6},5,2},
		{ []int{1,3,5,6},2,1},
		{ []int{1,3,5,6},7,4},
		{ []int{1,3,5,6},0,0},
		{ []int{1},1,0},
	}
	for i, t := range tests {
		var real = searchInsert(t.Nums, t.Target)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
