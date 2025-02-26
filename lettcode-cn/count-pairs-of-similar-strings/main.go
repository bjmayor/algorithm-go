package main

import "fmt"

func similarPairs(words []string) int {
	words2 := make([][26]int, len(words))
	n := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			words2[i][words[i][j]-'a'] = 1
		}
		for k := 0; k < i; k++ {
			if words2[k] == words2[i] {
				n++
			}
		}
	}
	return n
}

func main() {
	words := []string{"aba", "aabb", "abcd", "bac", "aabc"}
	fmt.Println(similarPairs(words))
}
