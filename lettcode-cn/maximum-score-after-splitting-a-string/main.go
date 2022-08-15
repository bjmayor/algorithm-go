package main

import "strings"

func maxScore(s string) int {
	sum := strings.Count(s,"1")
	result := 0
	for i:=0;i<len(s)-1;i++ {
		if s[i] == '0' {
			sum++
		} else {
			sum--
		}
		if  sum > result {
			result = sum
		}
	}
	return result
}

func main()  {
	println(maxScore("011101")) //5
	println(maxScore("00111")) //5
}