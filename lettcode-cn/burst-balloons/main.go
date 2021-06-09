package main

import (
	"fmt"
	"math"
)

func maxCoins(nums []int) int {
	n := len(nums)
	v := []int{1} //0,1....n,n+1
	v = append(v, nums...)
	v = append(v, 1)
	dp := make([][]int ,n+2)
	for i:=0;i<n+2;i++ {
		dp[i] = make([]int, n+2)
	}
	for i:=n-1;i>=0;i--{
		for j:=i+2;j<n+2;j++ {
			max := math.MinInt32
			for k:=i+1;k<=j-1;k++ {
				t := dp[i][k]	+ v[i]*v[k]*v[j] + dp[k][j]
				if t > max {
					max = t
				}
			}
			dp[i][j] =  max
		}
	}
	return dp[0][n+1]
}



func main()  {
	fmt.Println(maxCoins([]int{3})) //3
	fmt.Println(maxCoins([]int{3,1})) //6
	fmt.Println(maxCoins([]int{3,1,5})) //35
	fmt.Println(maxCoins([]int{3,1,5,8})) //167
}