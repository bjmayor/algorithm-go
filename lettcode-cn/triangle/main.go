package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	minPath := triangle[len(triangle)-1]//到达最底层的最小值
	for i:=len(triangle)-2;i>=0;i-- {
		for j:=0;j<=i;j++ {
				minPath[j] = triangle[i][j] + min(minPath[j], minPath[j+1])
		}
	}
	return minPath[0]
}

func min(x,y int) int {
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