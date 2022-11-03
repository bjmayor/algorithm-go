package main

import "strings"

func maxRepeating(sequence string, word string) int {
	var ans int
	for i := 1; ; i++ {
		if !strings.Contains(sequence, strings.Repeat(word, i)) {
			break
		} else {
			ans = i
		}
	}
	return ans
}

func main() {
	println(maxRepeating("a", "a"))
}
