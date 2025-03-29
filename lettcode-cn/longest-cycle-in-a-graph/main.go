package main

import "fmt"

/*
2360. 图中的最长环
*/
func longestCycle(edges []int) int {
	n := len(edges)
	visited := make([]bool, n)
	maxLength := -1

	for i := 0; i < n; i++ {
		if !visited[i] {
			currentTime := make(map[int]int)
			u := i
			for {
				// 没有出边
				if u == -1 {
					break
				}
				// 已经访问过
				if visited[u] {
					break
				}
				// 当前路径，已经访问过, 有环，计算环的长度
				if _, exists := currentTime[u]; exists {
					cycleLength := len(currentTime) - currentTime[u]
					if cycleLength > maxLength {
						maxLength = cycleLength
					}
					break
				}
				currentTime[u] = len(currentTime)
				u = edges[u]
			}
			for node := range currentTime {
				visited[node] = true
			}
		}
	}

	return maxLength
}

func main() {
	fmt.Println(longestCycle([]int{3, 3, 4, 2, 3})) // 3
	fmt.Println(longestCycle([]int{2, -1, 3, 1}))   // -1
}
