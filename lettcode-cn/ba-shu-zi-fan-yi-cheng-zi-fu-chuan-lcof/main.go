package main

import "fmt"

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	if num < 100 {
		if num < 26 {
			return 2
		} else {
			return 1
		}
	}
	if num % 100 <= 25 && num % 100 >=10 {
		return translateNum(num/10)	 + translateNum(num/100)
	} else {
		return translateNum(num/10)
	}
}

func main()  {
	tests := []struct{
		Nums int
		Expected int
	} {
		{ 12,2},
		{ 122,3},
		{ 1225,5},
		{ 12258,5},
	}
	for _, t := range tests {
		var real = translateNum(t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}