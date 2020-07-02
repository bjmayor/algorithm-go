package main

import "fmt"

//考虑dp求解。dp[i][j] A[...i]和B[...j]的最长公共子数组
func findLength(A []int, B []int) int {
	var res int
	dp := make([][]int, len(A)+1)
	for i:=0;i<=len(A);i++ {
		dp[i] = make([]int, len(B)+1)
	}

	for i:=0;i<len(A);i++ {
		for j:=0;j<len(A);j++	{
			if A[i] == B[j]	 {
				dp[i+1][j+1] = dp[i][j]+1
				if dp[i+1][j+1] > res {
					res = dp[i+1][j+1]
				}
			}

		}
	}
	return res
}

func main()  {
	fmt.Println(findLength([]int{1,2,3,2,1}, []int{3,2,1,4,7}))
}
