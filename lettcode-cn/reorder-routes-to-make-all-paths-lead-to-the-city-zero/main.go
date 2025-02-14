package main

import "fmt"

/**
 * minReorder 函数用于计算重新定向边的最小数量，使得所有路径都能通向城市0。
 * 算法思想：使用深度优先搜索（DFS）遍历图，统计需要重新定向的边的数量。
 *
 * @param n 整数，表示城市的数量
 * @param connections 二维整数数组，表示城市之间的连接
 * @return int 返回需要重新定向的边的最小数量
 */
func minReorder(n int, connections [][]int) int {
	g := make([][][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([][]int, 0)
	}
	for _, connection := range connections {
		x, y := connection[0], connection[1]
		g[x] = append(g[x], []int{y, 1})
		g[y] = append(g[y], []int{x, 0})
	}

	var dfs func(cur, parent int) int
	dfs = func(cur, parent int) int {
		var ans = 0
		for _, path := range g[cur] {
			if path[0] != parent {
				ans += dfs(path[0], cur)
				ans += path[1]
			}
		}
		return ans
	}
	return dfs(0, -1)
}

func main() {
	// 测试用例
	n := 6
	connections := [][]int{
		{0, 1},
		{1, 3},
		{2, 3},
		{4, 0},
		{4, 5},
	}
	// 调用minReorder函数并输出结果
	fmt.Println("需要重新定向的最小边数:", minReorder(n, connections)) // 输出: 3
}
