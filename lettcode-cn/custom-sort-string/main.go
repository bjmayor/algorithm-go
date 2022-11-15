package main

import "sort"

func customSortString(order string, s string) string {
	m := map[uint8]int{}
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}
	var bytes = []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return m[bytes[i]] < m[bytes[j]]
	})
	return string(bytes)
}

func main() {
	println(customSortString("kqep", "pekeq"))
}
