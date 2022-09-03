package main

import (
	"fmt"
)

func grayCode(n int) []int {
	max := 1 << n
	var ans = make([]int, max)
	for j, _ := range ans[1:] {
		ans[j+1] = (j + 1) ^ (j+1)>>1
	}
	return ans
}

func main() {
	fmt.Println(grayCode(3))
}
