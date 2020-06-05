package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	var sorted = make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	i :=0
	j:=len(sorted)-1

	var num1,num2 int
	for true {
		sum := sorted[i]+sorted[j]
		if  sum == target {
			num1 = sorted[i]
			num2 = sorted[j]
			break
		}	else if sum > target {
			j--
		} else {
			i++
		}
	}
	var result  = make([]int, 2)
	result[0] = -1
	result[1] = -1
	for i:=0;i<len(nums);i++ {
		if nums[i] == num1 && result[0]==-1{
			result[0] = i
			//continue;
		}
		if nums[i] == num2 && result[1]==-1 && i!=result[0] {
			result[1] = i
		}
	}
	return result
}

func arrEqual(a, b []int) bool  {
	if len(a) != len(b) {
		return false
	}
	for i:=0;i<len(a);i++ {
		if a[i]!=b[i] {
			return false
		}
	}
	return true
}

func main() {
	tests := []struct{
		Nums []int
		Target int
		Expected []int
	} {
		{ []int{2, 7, 11, 15}, 9, []int{0,1} },
		{ []int{3, 2, 4}, 6, []int{1,2} },
		{ []int{3, 3}, 6, []int{0,1} },
		{ []int{0,4,3,0}, 0, []int{0,3} },
		{ []int{-3,4,3,90}, 0, []int{0,2} },
		{ []int{-1,-2,-3,-4,-5}, -8, []int{2,4} },
	}
	for _, t := range tests {
		if !arrEqual(twoSum(t.Nums,t.Target), t.Expected) {
			fmt.Println("expected:", t.Expected, " real:", twoSum(t.Nums, t.Target))
		}
	}
}
