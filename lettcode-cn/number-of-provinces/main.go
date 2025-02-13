package main

import "fmt"

/*
https://leetcode.cn/problems/number-of-provinces/submissions/597900033/?envType=daily-question&envId=2025-02-09

547. 省份数量

有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
*/

/*
* findCircleNum 计算省份数量
* 算法思想:
* 1. 使用染色法(DFS)解决图的连通分量问题:
*    - visited 数组记录每个城市是否被访问过(染色)
*    - 每次从一个未访问的城市开始,通过 DFS 将所有与其相连的城市标记为已访问
* 2. 使用 DFS 遍历每个城市,统计连通分量(省份)的数量
* 3. 时间复杂度: O(n^2)
* 4. 空间复杂度: O(n)
 */

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	count := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(isConnected, visited, i)
			count++
		}
	}
	return count
}

func dfs(isConnected [][]int, visited []bool, i int) {
	visited[i] = true
	for j := 0; j < len(isConnected); j++ {
		if isConnected[i][j] == 1 && !visited[j] {
			dfs(isConnected, visited, j)
		}
	}
}

func main() {
	isConnected := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNum(isConnected))
}
