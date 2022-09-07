package main

import "strings"

func reorderSpaces(text string) string {
	words := strings.Fields(text)
	space := strings.Count(text, " ")
	size := len(words) - 1
	if size == 0 {
		return words[0] + strings.Repeat(" ", space)
	}
	return strings.Join(words, strings.Repeat(" ", space/size)) + strings.Repeat(" ", space%size)
}

func main() {
	println(reorderSpaces("  this   is  a sentence "))
	println(reorderSpaces("a"))
}
