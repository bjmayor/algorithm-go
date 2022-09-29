package main

func CheckPermutation(s1 string, s2 string) bool {
	m := map[rune]int{}
	for _, v := range s1 {
		m[v]++
	}
	for _, v := range s2 {
		m[v]--
		if m[v] == 0 {
			delete(m, v)
		}
	}
	return len(m) == 0
}
