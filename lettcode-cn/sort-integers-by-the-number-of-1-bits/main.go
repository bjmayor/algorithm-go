package main

import "sort"

func sortByBits(arr []int) []int {
	m := map[int]int{}
	for _, v := range arr {
		m[v] = numOfBits(v)
	}
	sort.Slice(arr, func(i, j int) bool {
		p, q := arr[i], arr[j]
		return m[p] < m[q] || m[p] == m[q] && p < q
	})
	return arr
}

func numOfBits(v int) int {
	var ans int
	for v > 0 {
		ans++
		v = v & (v - 1)
	}
	return ans
}
