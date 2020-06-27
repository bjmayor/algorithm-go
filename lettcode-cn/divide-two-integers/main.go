package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	var result int
	var sign = 1
	if dividend > 0 {
		if divisor < 0 {
			sign = -1
			divisor = -divisor
		}
	} else {
		dividend = -dividend
		if divisor > 0 {
			sign = -1
		} else {
			divisor = -divisor
		}
	}
	result = helper(dividend, divisor)
	if sign == -1 {
		result = -result
	}
	if result>math.MaxInt32 {
		result = math.MaxInt32
	}
	return result
}

func helper(dividend int, divisor int) int {
	if divisor > dividend {
		return 0
	}
	n := divisor
	i := 1
	for n <= dividend {
		n = n << 1
		i = i << 1
	}
	n = n >> 1
	i = i >> 1
	return i + helper(dividend-n, divisor)
}

func main() {
	fmt.Println(divide(10, 3)) //3
	fmt.Println(divide(-10,3))//-3
	fmt.Println(divide(-10,-3))//3
	fmt.Println(divide(10,-3))//-3
	fmt.Println(divide(7,-3))//-2
	fmt.Println(divide(1,1))//1
	fmt.Println(divide(math.MinInt32,1))//1
	fmt.Println(divide(math.MinInt32,-1))//2147483647
}
