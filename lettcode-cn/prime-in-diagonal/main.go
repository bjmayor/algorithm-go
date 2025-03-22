package main

import "fmt"

func diagonalPrime(nums [][]int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		if isPrime(nums[i][i]) {
			ans = max(ans, nums[i][i])
		}
		if isPrime(nums[i][len(nums)-1-i]) {
			ans = max(ans, nums[i][len(nums)-1-i])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(diagonalPrime([][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}}))
}
