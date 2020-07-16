package main

import "fmt"

var visited []bool
var colors []int
func isBipartite(graph [][]int) bool {
	n := len(graph)
	visited = make([]bool, n)
	colors = make([]int, n)
	for i:=0;i<n;i++ {
		if !visited[i] {
			if !dfs(i, 1, graph) {
				return false
			}
		}
	}
	return true
}

func dfs(idx int, color int, graph [][]int) bool  {
	visited[idx] = true
	colors[idx] = color
	oColor := -color
	for j:=0;j<len(graph[idx]);j++ {
		i := graph[idx][j]
		if !visited[i] {
			if !dfs(i, oColor, graph) {
				return false
			}
		} else if colors[i] != oColor {
			return false
		}
	}
	return true
}

func main()  {
fmt.Println(isBipartite([][]int{
	{1,3},
	{0,2},
	{1,3},
	{0,2},
	}))//true

	fmt.Println(isBipartite([][]int{
		{1,2,3},
		{0,2},
		{0,1,3},
		{0,2},
	}))//false

	fmt.Println(isBipartite([][]int{
		{1},
		{0,3},
		{3},
		{1,2},
	}))//true
}