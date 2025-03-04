package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	for _, v := range magazine {
		m[v-'a']++
	}
	for _, v := range ransomNote {
		m[v-'a']--
	}

	for _, v := range m {
		if v < 0 {
			return false
		}
	}
	return true
}

func main() {
	ransomNote := "a"
	magazine := "b"
	fmt.Println(canConstruct(ransomNote, magazine))
}
