package main

import "fmt"

func decrypt(code []int, k int) []int {
	var n = len(code)
	var ans []int = make([]int, n, n)
	if k == 0 {
		return ans
	}
	var s = 0
	if k < 0 {
		for i := n - 1; i >= n+k; i-- {
			s += code[i]
		}
	} else {
		for i := 1; i <= k; i++ {
			s += code[i]
		}
	}
	ans[0] = s
	for i, v := range code {
		if i > 0 {
			if k > 0 {
				s = s - v + code[(i+k)%n]
				ans[i] = s
			} else {
				s = s + code[(i-1+n)%n] - code[(i-1+k+n)%n]
				ans[i] = s
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))
	fmt.Println(decrypt([]int{1, 2, 3, 4}, 0))
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2))
}
