package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	ans := [][]int{}
	target := len(graph) - 1

	dfs(graph, 0, []int{0}, &ans, target)
	return ans
}

func dfs(graph [][]int, node int, path []int, ans *[][]int, target int) {
	if node == target {
		*ans = append(*ans, append([]int{}, path...))
		return
	}

	for _, neighbor := range graph[node] {
		dfs(graph, neighbor, append(path, neighbor), ans, target)
	}
}

func main() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println(allPathsSourceTarget(graph))
	graph = [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println(allPathsSourceTarget(graph))
	graph = [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}
