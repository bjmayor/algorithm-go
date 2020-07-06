package main

import "fmt"

func numWays(n int) int {
	if n ==0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	llast, last := 1,1
	for i:=2;i<=n;i++ {
		llast,	last = last, (llast+last)%1000000007
	}
	return last
}

func main() {
	tests := []struct{
		Nums int
		Expected int
	} {
		{ 2,2},
		{ 3,3},
		{ 44,134903163},
		{ 92,720754435},
	}
	for _, t := range tests {
		var real = numWays(t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}