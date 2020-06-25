package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	meet := false
	sign := 1
	val := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' && !meet {
			continue
		} else {
			c := str[i]
			if meet {
				if !expected(c, '0','1','2','3','4','5','6','7','8','9') {
					break
				}
			} else {
				meet = true
				if !expected(c, '+','-','0','1','2','3','4','5','6','7','8','9') {
					break
				} else {
					if c == '-' {
						sign = -1
						c = '0'
					} else if c == '+' {
						c = '0'
					}
				}
			}
			val = val*10+int(c-'0')
			if val > math.MaxInt32 {
				break
			}
		}
	}

	val = sign * val
	if val > math.MaxInt32 {
		val = math.MaxInt32
	} else if val < math.MinInt32 {
		val = math.MinInt32
	}
	return val
}

func expected(v byte, values ...byte) bool {
	for _, item := range values {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	tests := []struct {
		s        string
		Expected int
	}{
		//{"42", 42},
		{"+-2", 0},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"9223372036854775808", 2147483647},
	}
	for i, t := range tests {
		var real = myAtoi(t.s)
		if real != t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real)
		}
	}
}

