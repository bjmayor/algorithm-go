package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if  x %10 == 0 {
		return false
	}
	var r int
	for r<x {
		r = r * 10 + x%10
		x = x/10
	}
	return r==x || r/10 == x
}

func main()  {
	tests := []struct{
		Nums int
		Expected bool
	} {
		{ 1001,true},
	}
	for _, t := range tests {
		var real = isPalindrome(t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}