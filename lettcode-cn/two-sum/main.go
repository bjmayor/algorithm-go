package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i,v := range nums {
		m[v] = i
	}
	for i,v := range nums {
		if vv, ok := m[target-v];ok && i!=vv {
			return []int{i,vv}
		}
	}
	return nil
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
		//{ []int{2, 7, 11, 15}, 9, []int{0,1} },
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
