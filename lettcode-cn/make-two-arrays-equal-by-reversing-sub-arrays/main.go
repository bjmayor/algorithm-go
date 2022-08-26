package main

func canBeEqual(target []int, arr []int) bool {
	m := map[int]int{}
	if len(target) != len(arr) {
		return false
	}
	for i, v := range target {
		m[v]++
		m[arr[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
