package main

import "fmt"

func countBits(num int) []int {
	var result = make([]int, num+1, num+1)
	for i := 1; i <= num; i++ {
		if i&(i-1) == 0 {
			for j := 0; j < i && i+j <= num; j++ {
				result[i+j] = 1 + result[j]
			}
		}
	}
	return result
}


func countBits2(num int) []int {
	var result = make([]int, num+1, num+1)
	for i := 1; i <= num; i++ {
		if i&(i-1) == 0 {
			for j := 0; j < i && i+j <= num; j++ {
				result[i+j] = 1 + result[j]
			}
		}
	}
	return result
}


func main() {
	fmt.Println(countBits2(5))
}
