package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums)-1
	for left <= right{
		if right>left {
			rnd := rand.Intn(right-left)+left
			nums[rnd], nums[left] = nums[left], nums[rnd]
		} else {
			return nums[left]
		}

		key := nums[left]
		i:=left+1 //i 左边都比k大
		d := right

		for i<=d {
			v := nums[i]
			if  v < key {
				nums[i], nums[d]	 = nums[d], nums[i]
				d--
			} else {
				i++
			}
			if d >=i && nums[d] > key {
				nums[i], nums[d]	 = nums[d], nums[i]
				i++
			}
		}
		nums[left], nums[i-1] = nums[i-1], nums[left]
		if i == k {
			return key
		} else if i > k {
			right = i-2
		} else {
			left = i
		}
	}
	return 0
}

func main()  {
	tests := []struct{
		Nums []int
		Target int
		Expected int
	} {
		//{ []int{3,2,1,5,6,4},2,5},
		{ []int{3,2,3,1,2,4,5,5,6},4,4},
	}
	for i, t := range tests {
		var real = findKthLargest(t.Nums, t.Target)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}