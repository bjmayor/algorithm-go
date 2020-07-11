package main

import "fmt"

func isSubsequence(s string, t string) bool {
	i := 0;
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	if i == len(s) {
		return true
	}
	return false
}

func main() {
	tests := []struct {
		s        string
		t        string
		Expected bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for i, t := range tests {
		var real = isSubsequence(t.s, t.t)
		if real != t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real)
		}
	}
}
