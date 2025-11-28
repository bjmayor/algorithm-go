package main

import "fmt"

func minOperations(nums []int) int {
	n := len(nums)

	// 统计1的个数
	countOnes := 0
	for _, num := range nums {
		if num == 1 {
			countOnes++
		}
	}

	// 情况1：已经有1
	if countOnes > 0 {
		return n - countOnes
	}

	// 情况2：没有1，检查整体gcd
	overallGcd := nums[0]
	for i := 1; i < n; i++ {
		overallGcd = gcd(overallGcd, nums[i])
	}

	if overallGcd > 1 {
		return -1 // 情况3：不可能
	}

	// 情况4：寻找最短子数组
	minLen := n + 1
	for i := 0; i < n; i++ {
		currentGcd := nums[i]
		for j := i + 1; j < n; j++ {
			currentGcd = gcd(currentGcd, nums[j])
			if currentGcd == 1 {
				minLen = min(minLen, j-i+1)
				break
			}
		}
	}

	return n + minLen - 2
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 测试用例
	testCases := [][]int{
		{2, 6, 3, 4},     // 期望: 4
		{2, 10, 3, 9, 6}, // 期望: 5
		{1, 1, 1, 1},     // 期望: 0
		{6, 10, 15},      // 期望: 4
		{2, 4, 6, 8},     // 期望: -1
	}

	for _, nums := range testCases {
		result := minOperations(nums)
		fmt.Printf("nums: %v, result: %d\n", nums, result)
	}
}
