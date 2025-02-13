package main

import "fmt"

/*
https://leetcode.cn/problems/find-eventual-safe-states/description/?envType=study-plan-v2&envId=graph-theory

解题思路:
1. 使用深度优先搜索(DFS)来判断每个节点是否是安全的
2. 使用一个状态数组来记录每个节点的状态
3. 状态数组有三种状态:
   - 0: 未访问
   - 1: 访问中
   - 2: 安全
4. 从每个未访问的节点开始进行DFS搜索
5. 如果搜索到一个访问中的节点,说明存在环,该节点不是安全节点
6. 如果搜索到一个安全的节点,则该节点是安全节点
7. 返回所有安全节点的列表
*/

func eventualSafeNodes(graph [][]int) []int {
	// 状态数组,0表示未访问,1表示访问中,2表示安全
	status := make([]int, len(graph))
	result := []int{}
	// 遍历每个节点
	for i := 0; i < len(graph); i++ {
		if dfs(graph, status, i) {
			result = append(result, i)
		}
	}
	return result
}

func dfs(graph [][]int, status []int, i int) bool {
	// 如果节点已经被访问过,并且是访问中的节点,说明存在环,该节点不是安全节点
	if status[i] == 1 {
		return false
	}
	// 如果节点已经被访问过,并且是安全的节点,则该节点是安全节点
	if status[i] == 2 {
		return true
	}
	// 将当前节点标记为访问中
	status[i] = 1
	// 遍历当前节点的所有邻居节点
	for _, neighbor := range graph[i] {
		// 如果邻居节点不是安全节点,则当前节点也不是安全节点
		if !dfs(graph, status, neighbor) {
			return false
		}
	}
	status[i] = 2
	return true
}

func main() {
	graph := [][]int{
		{1, 2},
		{2, 3},
		{5},
		{0},
		{5},
		{},
		{},
	}
	fmt.Println(eventualSafeNodes(graph))
}
