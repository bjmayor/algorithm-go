package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	numStr := strconv.Itoa(num)
	count := 0
	for i := 0; i <= len(numStr)-k; i++ {
		substring := numStr[i : i+k]
		substringInt, _ := strconv.Atoi(substring)
		if substringInt > 0 && num%substringInt == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(divisorSubstrings(240, 2))
}
