package main

import (
	"fmt"
	"math"
)

func palindromePartition(s string, k int) int {
	n := len(s)
	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
	}
	for span := 2; span <= n; span++ {
		for i := 0; i <= n-span; i++ {
			j := i + span - 1
			cost[i][j] = cost[i+1][j-1]
			if s[i] != s[j] {
				cost[i][j]++
			}
		}
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt
		}
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(k, i); j++ {
			if j == 1 {
				f[i][j] = cost[0][i-1]
			} else {
				for i0 := j - 1; i0 < i; i0++ {
					f[i][j] = min(f[i][j], f[i0][j-1]+cost[i0][i-1])
				}
			}
		}
	}

	return f[n][k]
}

func main() {
	fmt.Println(palindromePartition("abc", 2))
}
