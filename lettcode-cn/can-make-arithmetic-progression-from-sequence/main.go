package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	v := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] != v {
			return false
		}
	}
	return true
}
