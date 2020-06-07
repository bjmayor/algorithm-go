package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var y int
	for x !=0 {
		y = y*10
		y += x%10
		x = x/10
		if y > math.MaxInt32 || y < math.MinInt32 {
			return 0
		}
	}

	return y
}


func main()  {
	tests := []struct{
		Nums int
		Expected int
	} {
		{ 10,1},
		{ 123,321},
		{ -123,-321},
		{ 120,21},
		{ 1534236469,0},
	}
	for _, t := range tests {
		var real = reverse(t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}