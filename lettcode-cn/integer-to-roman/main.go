package main

import "fmt"

func intToRoman(num int) string {
	var result []byte
	n :=  num/1000
	for i:=0;i<n;i++ {
		result = append(result, 'M')
	}
	num = num%1000
	if num >= 900 {
		result = append(result, 'C','M')
		num -= 900
	}
	if num >= 500 {
		result = append(result, 'D')
		num -= 500
	}

	if num >= 400 {
		result = append(result, 'C','D')
		num -= 400
	}
	n =  num/100
	for i:=0;i<n;i++ {
		result = append(result, 'C')
	}
	num = num % 100

	if num >= 90 {
		result = append(result, 'X','C')
		num -= 90
	}
	if num >= 50 {
		result = append(result, 'L')
		num -= 50
	}

	if num >= 40 {
		result = append(result, 'X','L')
		num -= 40
	}
	n =  num/10
	for i:=0;i<n;i++ {
		result = append(result, 'X')
	}
	num = num % 10
	if num >= 9 {
		result = append(result, 'I','X')
		num -= 9
	}
	if num >= 5 {
		result = append(result, 'V')
		num -= 5
	}

	if num >= 4 {
		result = append(result, 'I','V')
		num -= 4
	}
	for i:=0;i<num;i++ {
		result = append(result, 'I')
	}

	return string(result)
}

func main()  {
	tests := []struct {
		s       int
		Expected string
	}{
		{3,"III"},
		{4,"IV"},
		{9,"IX"},
		{58,"LVIII"},
		{1994,"MCMXCIV"},
	}
	for i, t := range tests {
		var real = intToRoman(t.s)
		if real != t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real)
		}
	}
}