package main

import "fmt"

func addSpaces(s string, spaces []int) string {
	ans := make([]byte, 0, len(s)+len(spaces))
	last := 0
	for _, space := range spaces {
		ans = append(ans, s[last:space]...)
		ans = append(ans, ' ')
		last = space
	}
	ans = append(ans, s[last:]...)
	return string(ans)
}

func main() {
	fmt.Println(addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15})) // "Leetcode Helps Me Learn"
}
