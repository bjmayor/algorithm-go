package main

import (
	"fmt"
)

func longestIncreasingPath(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i:=0;i<len(matrix);i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	var res = 0
	for i:=0;i<len(matrix);i++ {
		for j:=0;j<len(matrix[i]);j++ {
			if dp[i][j] == 0 {
				tmp := dfs(matrix, dp, i, j)
				if tmp > res {
					res = tmp
				}
			}
		}
	}
	return res
}

func dfs(matrix [][]int, dp [][]int, i,j int) int  {
	if dp[i][j] == 0 {
		var res = 1
		if i-1>=0 && matrix[i-1][j]>matrix[i][j] {
			o := dfs(matrix, dp, i-1,j)
			if 1+o > res {
				res = 1+o
			}
		}
		if i+1 < len(matrix) && matrix[i+1][j]>matrix[i][j] {
			o := dfs(matrix, dp, i+1,j)
			if 1+o > res {
				res = 1+o
			}
		}

		if j-1>=0 && matrix[i][j-1]>matrix[i][j] {
			o := dfs(matrix, dp, i,j-1)
			if 1+o > res {
				res = 1+o
			}
		}

		if j+1 < len(matrix[i]) && matrix[i][j+1]>matrix[i][j] {
			o := dfs(matrix, dp, i,j+1)
			if 1+o > res {
				res = 1+o
			}
		}
		dp[i][j] = res
	}

	 return dp[i][j]
}
func main()  {
	tests := []struct{
		Nums [][]int
		Expected int
	} {
		{ [][]int{
			{9,9,4},
			{6,6,8},
			{2,2,1},
		} ,4 },
		{ [][]int{
			{3,4,5},
			{3,2,6},
			{2,2,1},
		} ,4 },
	}
	for i, t := range tests {
		v := longestIncreasingPath(t.Nums)
		if v !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", v)
		}
	}
}
