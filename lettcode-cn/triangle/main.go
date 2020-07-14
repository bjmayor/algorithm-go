package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	minPath := triangle[len(triangle)-1]
	for j:=len(triangle)-2;j>=0;j-- {
		for i:=0;i<=j;i++ {
			minPath[i] = triangle[j][i]+min(minPath[i], minPath[i+1])
		}
	}
	return minPath[0]
}

func min(x,y int) int  {
	if x < y {
		return x
	}
	return y
}


func main()  {
	triangle:=[][]int{
		{2},
		{3,4},
		{6,5,7},
		{4,1,8,3},
	}
	fmt.Println(minimumTotal(triangle))
}