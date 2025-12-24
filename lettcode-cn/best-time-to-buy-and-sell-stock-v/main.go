package main

import (
	"fmt"
)

func maximumProfit(prices []int, k int) int64 {
	n := len(prices)
	dp := make([][][]int64, n)
	for i := range dp {
		dp[i] = make([][]int64, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, 3)
		}
	}

	// 初始化第 0 天的状态
	for j := 1; j <= k; j++ {
		dp[0][j][1] = int64(-prices[0])
		dp[0][j][2] = int64(prices[0])
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], max(dp[i-1][j][1]+int64(prices[i]), dp[i-1][j][2]-int64(prices[i])))
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-int64(prices[i]))
			dp[i][j][2] = max(dp[i-1][j][2], dp[i-1][j-1][0]+int64(prices[i]))
		}
	}

	return dp[n-1][k][0]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 示例 1
	prices1 := []int{1, 7, 9, 8, 2}
	k1 := 2
	fmt.Printf("示例1: prices=%v, k=%d, 结果=%d\n", prices1, k1, maximumProfit(prices1, k1))

	// 示例 2
	prices2 := []int{12, 16, 19, 19, 8, 1, 19, 13, 9}
	k2 := 3
	fmt.Printf("示例2: prices=%v, k=%d, 结果=%d\n", prices2, k2, maximumProfit(prices2, k2))

	// 之前失败的测试用例
	prices3 := []int{8, 4, 15, 7, 4, 7, 2, 14, 15}
	k3 := 3
	fmt.Printf("测试用例: prices=%v, k=%d, 结果=%d (预期: 28)\n", prices3, k3, maximumProfit(prices3, k3))
}
