package main

import "fmt"

/*
https://leetcode.cn/problems/shortest-path-with-alternating-colors/description/?envType=study-plan-v2&envId=graph-theory
*/

/**
 * shortestAlternatingPaths 函数计算从节点 0 到其他所有节点的最短路径，路径中相邻边的颜色必须交替。
 * @param n 节点的数量
 * @param redEdges 红色边的列表
 * @param blueEdges 蓝色边的列表
 * @return 返回一个整数数组，表示从节点 0 到每个节点的最短路径长度，如果无法到达则为 -1
 */
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	// 初始化结果数组，所有节点的初始距离为 -1
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}
	result[0] = 0 // 起始节点到自身的距离为 0

	// 构建图的邻接表
	redGraph := make(map[int][]int)
	blueGraph := make(map[int][]int)
	for _, edge := range redEdges {
		redGraph[edge[0]] = append(redGraph[edge[0]], edge[1])
	}
	for _, edge := range blueEdges {
		blueGraph[edge[0]] = append(blueGraph[edge[0]], edge[1])
	}

	// 使用队列进行广度优先搜索
	type state struct {
		node  int
		color int // 0 表示红色，1 表示蓝色
	}
	queue := []state{{0, 0}, {0, 1}} // 从节点 0 开始，尝试两种颜色
	visited := make(map[state]bool)
	visited[state{0, 0}] = true
	visited[state{0, 1}] = true

	steps := 0
	for len(queue) > 0 {
		steps++
		nextQueue := []state{}
		for _, s := range queue {
			var neighbors []int
			if s.color == 0 {
				neighbors = blueGraph[s.node]
			} else {
				neighbors = redGraph[s.node]
			}
			for _, neighbor := range neighbors {
				nextState := state{neighbor, 1 - s.color}
				if !visited[nextState] {
					visited[nextState] = true
					nextQueue = append(nextQueue, nextState)
					if result[neighbor] == -1 {
						result[neighbor] = steps
					}
				}
			}
		}
		queue = nextQueue
	}

	return result
}

func main() {
	// 测试用例1: n = 3, redEdges = [[0,1],[1,2]], blueEdges = []
	n1 := 3
	redEdges1 := [][]int{{0, 1}, {1, 2}}
	blueEdges1 := [][]int{}
	result1 := shortestAlternatingPaths(n1, redEdges1, blueEdges1)
	fmt.Println("测试用例1结果:", result1) // 预期输出: [0, 1, -1]

	// 测试用例2: n = 3, redEdges = [[0,1]], blueEdges = [[2,1]]
	n2 := 3
	redEdges2 := [][]int{{0, 1}}
	blueEdges2 := [][]int{{2, 1}}
	result2 := shortestAlternatingPaths(n2, redEdges2, blueEdges2)
	fmt.Println("测试用例2结果:", result2) // 预期输出: [0, 1, -1]

	// 测试用例3: n = 3, redEdges = [[0,1]], blueEdges = [[1,2]]
	n3 := 3
	redEdges3 := [][]int{{0, 1}}
	blueEdges3 := [][]int{{1, 2}}
	result3 := shortestAlternatingPaths(n3, redEdges3, blueEdges3)
	fmt.Println("测试用例3结果:", result3) // 预期输出: [0, 1, 2]
}
