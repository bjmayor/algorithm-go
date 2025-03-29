package main

import (
	"fmt"
	"strings"
)

func countPrefixes(words []string, s string) int {
	count := 0
	for _, word := range words {
		if strings.HasPrefix(s, word) {
			count++
		}
	}
	return count
}

func main() {
	words := []string{"a", "b", "c", "ab", "bc", "abc"}
	s := "abc"
	fmt.Println(countPrefixes(words, s))
}
