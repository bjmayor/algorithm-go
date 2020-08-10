package main

import "fmt"

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)
	for j:=n-1;j>=0;j-- {
		if j < n-1 {
			dp[j] = grid[m-1][j] + dp[j+1]
		} else {
			dp[j] = grid[m-1][j]
		}
	}
	for i:=m-2;i>=0;i-- {
		for j:= n-1;j>=0;j-- {
			if j < n-1 {
				dp[j] = grid[i][j]+min(dp[j+1], dp[j])
			} else {
				dp[j] = grid[i][j]+dp[j]
			}
		}
	}
	return dp[0]
}

func min(x,y int) int  {
	if x < y {
		return x
	}
	return y
}

func main()  {
fmt.Print(minPathSum([][]int{
	{1,3,1},
	{1,5,1},
	{4,2,1},
}))
}
