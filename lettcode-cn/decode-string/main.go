package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			tmp := []byte{}
			for stack[len(stack)-1] != '[' {
				tmp = append([]byte{stack[len(stack)-1]}, tmp...)
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			t := string(tmp)
			tmp = []byte{}
			top := stack[len(stack)-1]
			for top >= '0' && top <= '9' {
				tmp = append([]byte{stack[len(stack)-1]}, tmp...)
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					break
				}
				top = stack[len(stack)-1]
			}
			n, _ := strconv.Atoi(string(tmp))
			origin := []byte(strings.Repeat(t, n))
			stack = append(stack, origin...)
		} else {
			stack = append(stack, s[i]) // 入栈
		}
	}
	return string(stack)
}

func main() {
	//println(decodeString("3[a]2[bc]"))
	println(decodeString("100[leetcode]"))
}
