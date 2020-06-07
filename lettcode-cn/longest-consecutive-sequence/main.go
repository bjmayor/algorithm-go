package main

import "fmt"


func longestConsecutive(nums []int) int {
	m := make(map[int]struct{},0)
	var result = 0
	if len(nums) > 0 {
		result = 1
	}
	for _, k := range nums {
		m[k] = struct{}{}
	}
	for k,_ := range m {
		if _, ok := m[k-1]; ok {
			continue //不是最长子序列的起始
		}
		len := 1
		detect := k+1
		//不断试探
		for {
			if _, ok := m[detect]; !ok {
				break
			}
			len ++
			detect++
		}
		if len > result {
			result = len
		}
	}

	return result
}



func main() {
	tests := []struct {
		Nums     []int
		Expected int
	}{
		//{[]int{100, 4, 200, 1, 3, 2}, 4},
		//{[]int{0}, 1},
		{[]int{1,3,5,2,4}, 5},
	}
	for _, t := range tests {
		var real = longestConsecutive(t.Nums)
		if real != t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real)
		}
	}
}
