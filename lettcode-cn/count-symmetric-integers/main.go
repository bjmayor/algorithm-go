package main

import (
	"fmt"
)

/*
2843. 统计对称整数的数目
给你两个正整数 low 和 high 。

对于一个由 2 * n 位数字组成的整数 x ，如果其前 n 位数字之和与后 n 位数字之和相等，则认为这个数字是一个对称整数。

返回在 [low, high] 范围内的 对称整数的数目 。
*/
func countSymmetricIntegers(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		if isSymmetric(i) {
			count++
		}
	}
	return count
}

func isSymmetric(num int) bool {
	// Check if the number is in the valid range for symmetric numbers
	if (num >= 11 && num <= 99) || (num >= 1000 && num <= 9999) {
		// Calculate the sum of the first half and the second half
		if num >= 11 && num <= 99 {
			// Two-digit number
			firstDigit := num / 10
			secondDigit := num % 10
			return firstDigit == secondDigit
		} else {
			// Four-digit number
			firstHalf := num / 100
			secondHalf := num % 100
			firstSum := firstHalf/10 + firstHalf%10
			secondSum := secondHalf/10 + secondHalf%10
			return firstSum == secondSum
		}
	}
	return false
}

func main() {
	fmt.Println(countSymmetricIntegers(1, 100))     // 9
	fmt.Println(countSymmetricIntegers(1200, 1230)) // 4
}
