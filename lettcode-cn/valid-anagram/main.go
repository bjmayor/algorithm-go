package main

import "fmt"

func isAnagram(s string, t string) bool {
	var a1 [26]int
	var a2 [26]int
	for i := 0; i < len(s); i++ {
		a1[s[i]-'a'] += 1
	}

	for i := 0; i < len(t); i++ {
		a2[t[i]-'a'] += 1
	}
	return a1 == a2
}

func main() {
	tests := []struct {
		s        string
		t        string
		Expected bool
	}{
		//{"anagram", "nagaram", true},
		//{"rat", "car", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbba", false},
	}
	for _, t := range tests {
		if isAnagram(t.s, t.t) != t.Expected {
			fmt.Println("expected:", t.Expected, " real:", isAnagram(t.s, t.t))
		} else {
			fmt.Println("ok")
		}
	}
}
