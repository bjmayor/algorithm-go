package main

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	return strings.Contains(s2+s2, s1)
}
