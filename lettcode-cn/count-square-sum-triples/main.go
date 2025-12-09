package main

import "math"

func countTriples(n int) int {
	// a<b<c<n
	cnt := 0
	for a := 1; a < n; a++ {
		for b := a + 1; b < n; b++ {
			c := a*a + b*b
			if c <= n*n {
				sqrtC := int(math.Sqrt(float64(c)))
				if sqrtC*sqrtC == c {
					cnt += 2
				}
			}
		}
	}
	return cnt
}

func main() {
	println(countTriples(5))  // expect 2
	println(countTriples(10)) // expect 4
}
