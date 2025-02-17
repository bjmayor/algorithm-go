package main

import "fmt"

func lengthOfLastWord(s string) int {
	result := 0
	l := 0
	for _, v := range s {
		if v == ' ' {
			if l > 0 {
				result = l
				l = 0
			}
			continue
		} else {
			l++
		}

	}
	if l > 0 {
		result = l
	}
	return result
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord(" "))
}
