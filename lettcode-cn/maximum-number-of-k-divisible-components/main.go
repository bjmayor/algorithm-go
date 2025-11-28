package main

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {

	// 建图：边列表转邻接表
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	ans := 0

	// DFS: 自底向上收集子树和
	var dfs func(node, parent int) int
	dfs = func(node, parent int) int {
		// 当前节点的初始和
		sum := values[node]

		// 收集所有子树的和
		for _, child := range graph[node] {
			if child == parent {
				continue // 避免回到父节点
			}
			childSum := dfs(child, node)
			sum += childSum
		}

		// 关键判断：如果当前子树和能被k整除
		if sum%k == 0 {
			ans++    // 切断与父节点的边，形成独立块
			return 0 // 向上传递0，表示这部分已经独立
		}

		return sum // 否则向上传递和
	}

	dfs(0, -1) // 从节点0开始，-1表示无父节点
	return ans
}
