package main // 定义包名为main

import "fmt" // 导入fmt包用于格式化I/O

/*
https://leetcode.cn/problems/critical-connections-in-a-network/description/?envType=study-plan-v2&envId=graph-theory

1192. 查找集群内的关键连接

解题思路：
1. 使用Tarjan算法来查找图中的桥。
2. 使用两个数组：
   - `visited` 记录每个节点是否被访问过。
   - `low` 记录每个节点通过DFS能访问到的最低节点。
   - `dfn` 记录每个节点被访问的时间戳。
3. 如果 `low[v] > dfn[u]`，则边 (u, v) 是关键连接。

*/

func criticalConnections(n int, connections [][]int) [][]int { // 定义函数criticalConnections，输入节点数n和连接数组connections，返回关键连接数组
	graph := make([][]int, n) // 初始化图的邻接表
	for _, connection := range connections { // 遍历每个连接
		graph[connection[0]] = append(graph[connection[0]], connection[1]) // 添加无向边
		graph[connection[1]] = append(graph[connection[1]], connection[0]) // 添加无向边
	}
	visited := make([]bool, n) // 初始化访问标记数组
	low := make([]int, n) // 初始化low数组
	dfn := make([]int, n) // 初始化dfn数组
	timestamp := 0 // 初始化时间戳
	ans := [][]int{} // 初始化答案数组

	var tarjan func(u, parent int) // 定义Tarjan算法的递归函数
	tarjan = func(u, parent int) { // 实现Tarjan算法
		visited[u] = true // 标记节点u为已访问
		timestamp++ // 增加时间戳
		dfn[u] = timestamp // 设置dfn[u]为当前时间戳
		low[u] = timestamp // 设置low[u]为当前时间戳
		for _, v := range graph[u] { // 遍历节点u的邻接节点
			if v == parent { // 如果v是父节点，跳过
				continue // 继续下一个邻接节点
			}
			if !visited[v] { // 如果v未被访问
				tarjan(v, u) // 递归访问v
				low[u] = min(low[u], low[v]) // 更新low[u]
				if low[v] > dfn[u] { // 如果low[v]大于dfn[u]
					ans = append(ans, []int{u, v}) // 将边(u, v)加入答案
				}
			} else { // 如果v已被访问
				low[u] = min(low[u], dfn[v]) // 更新low[u]
			}
		}
	}
	for i := 0; i < n; i++ { // 遍历所有节点
		if !visited[i] { // 如果节点i未被访问
			tarjan(i, -1) // 从节点i开始执行Tarjan算法
		}
	}

	return ans // 返回关键连接数组
}

func min(a, b int) int { // 定义函数min，返回两个整数中的较小值
	if a < b { // 如果a小于b
		return a // 返回a
	}
	return b // 否则返回b
}

func main() { // 主函数
	n := 4 // 定义节点数为4
	connections := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}} // 定义连接数组
	fmt.Println(criticalConnections(n, connections)) // 输出关键连接
}
