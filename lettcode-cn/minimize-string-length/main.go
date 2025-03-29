package main

import "fmt"

/*
2716. 最小化字符串长度
*/
func minimizedStringLength(s string) int {
	set := make(map[rune]struct{})
	for _, v := range s {
		set[v] = struct{}{}
	}
	return len(set)
}

func main() {
	fmt.Println(minimizedStringLength("aaabc"))
}
