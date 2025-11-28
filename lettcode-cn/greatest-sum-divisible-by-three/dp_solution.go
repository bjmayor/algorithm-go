package main

import (
	"fmt"
)

func maxSumDivThreeDP(nums []int) int {
	const INF = 1 << 30
	dp := [3]int{-INF, -INF, -INF}
	dp[0] = 0
	for _, num := range nums {
		mod := num % 3
		newDp := [3]int{dp[0], dp[1], dp[2]}
		for j := 0; j < 3; j++ {
			if dp[j] != -INF {
				newMod := (j + mod) % 3
				newDp[newMod] = max(newDp[newMod], dp[j]+num)
			}
		}
		dp = newDp
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 示例1
	nums1 := []int{3, 6, 5, 1, 8}
	fmt.Println(maxSumDivThreeDP(nums1)) // 18

	// 示例2
	nums2 := []int{4}
	fmt.Println(maxSumDivThreeDP(nums2)) // 0

	// 示例3
	nums3 := []int{1, 2, 3, 4, 4}
	fmt.Println(maxSumDivThreeDP(nums3)) // 12
}
