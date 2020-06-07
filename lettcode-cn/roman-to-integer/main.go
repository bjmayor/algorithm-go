package main

import (
	"fmt"
)

func romanToInt(s string) int {
	dict := map[rune]int{
		'I':1,
		'V':5,
		'X':10,
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}
	var result int
	var pre rune
	for _, c := range s{
		result += dict[c]
		if pre > 0 && dict[pre]<dict[c] {
			result -= dict[pre] * 2
		}
		pre = c
	}
	return result
}

func main()  {
	tests := []struct{
		Nums string
		Expected int
	} {
		{ "III" ,3 },
		{ "IV" ,4 },
		{ "IX" ,9 },
		{ "LVIII" ,58 },
		{ "MCMXCIV" ,1994 },
	}
	for _, t := range tests {
		if romanToInt(t.Nums) !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", romanToInt(t.Nums), t.Nums)
		}
	}
}
