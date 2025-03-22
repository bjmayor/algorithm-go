package main

import (
	"fmt"
	"math"
)

/*
给你一个字符串 s 。一个字符串的 分数 定义为相邻字符 ASCII 码差值绝对值的和。

请你返回 s 的 分数 。
*/
func scoreOfString(s string) int {
	score := 0
	for i := 0; i < len(s)-1; i++ {
		score += int(math.Abs(float64(s[i]) - float64(s[i+1])))
	}
	return score
}

func main() {
	fmt.Println(scoreOfString("hello")) // 13
	fmt.Println(scoreOfString("zaz"))   // 50
}
