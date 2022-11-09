package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   bool = true // 检测有没有环
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		//edges[info[1]] = append(edges[info[1]], info[0])
		edges[info[0]] = append(edges[info[0]], info[1])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	//for i := 0; i < len(result)/2; i++ {
	//	result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	//}
	return result
}

func main() {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(1, [][]int{}))
	fmt.Println(findOrder(2, [][]int{{0, 1}}))
}
