package main

import "fmt"

func constructArray(n int, k int) []int {
	var ans = make([]int, 0, n)
	// 分两部分。 第一部分 1,2,3,n-k+1
	for i := 1; i < n-k; i++ {
		ans = append(ans, i)
	}
	// 第二部分 n-k, n, n-k+1, n-1, ...
	for i, j := n-k, n; i <= j; i++ {
		ans = append(ans, i)
		if i != j {
			ans = append(ans, j)
		}
		j--
	}
	return ans
}

func main() {
	fmt.Println(constructArray(10, 3))
}
