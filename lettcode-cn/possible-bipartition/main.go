package main

func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n) // 整理下，把不喜欢的人 放一起
	for _, d := range dislikes {
		x, y := d[0]-1, d[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	color := make([]int, n) // color[x] = 0 表示未访问节点 x, 也设置颜色1,2, 3^c就是1->2,2->1
	var dfs func(int, int) bool
	dfs = func(x, c int) bool {
		color[x] = c // x 设置为c, 则 x不喜欢的y 如果已经设置成了c 就不符合要求。
		for _, y := range g[x] {
			if color[y] == c || color[y] == 0 && !dfs(y, 3^c) {
				return false
			}
		}
		return true
	}
	for i, c := range color {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}

func main() {
	//println(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
	//println(possibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	//println(possibleBipartition(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}))
	println(possibleBipartition(10, [][]int{{1, 2}, {3, 4}, {5, 6}, {6, 7}, {8, 9}, {7, 8}}))
}
