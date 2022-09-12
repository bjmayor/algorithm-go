package main

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diff := []int{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
		}
	}
	if len(diff) == 0 {
		return true
	}
	if len(diff) == 2 && s1[diff[0]] == s2[diff[1]] && s1[diff[1]] == s2[diff[0]] {
		return true
	}
	return false
}
