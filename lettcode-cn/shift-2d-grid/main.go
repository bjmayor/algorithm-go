package main

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	total := m * n
	result := make([][]int,m,m)
	for i:=0;i<m;i++ {
		result[i] = make([]int, n, n)
	}
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			idx := (i*n+j-k%total+total)%total
			result[i][j] = grid[idx/n][idx%n]
		}
	}
	return result
}

func main()  {
	//fmt.Println(shiftGrid([][]int{{1,2,3},{4,5,6},{7,8,9}}, 9))
	fmt.Println(shiftGrid([][]int{{1},{2},{3},{4},{7},{6},{5}}, 23))
}