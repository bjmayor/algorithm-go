package main

import (
	"fmt"
	"math"
)

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	dp := make([][]int, m) //可营救的最小值
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = max(1,1 - dungeon[i][j])
			}
			if i-1 >= 0  {
				dp[i-1][j] = min(max(1,dp[i][j]-dungeon[i-1][j]), dp[i-1][j])
			}
			if j-1 >= 0  {
				dp[i][j-1] = min(max(1,dp[i][j]-dungeon[i][j-1]), dp[i][j-1])
			}
		}

	}
	return dp[0][0]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int)  int {
	if x > y {
		return x
	}
	return y
}


func main() {
	fmt.Println(calculateMinimumHP([][]int{
		{-2, -3, 3,},
		{-5, -10, 1,},
		{10, 30, -5,},
	}))
}
