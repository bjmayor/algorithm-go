package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func pacificAtlantic(heights [][]int) (ans [][]int) {
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)  // 记录可以到达太平洋的点
	atlantic := make([][]bool, m) // 记录可以到达大西洋的点
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func(int, int, [][]bool)
	// 点 (x,y) 能否留到ocean
	dfs = func(x, y int, ocean [][]bool) {
		if ocean[x][y] {
			return
		}
		ocean[x][y] = true
		// 判断4个方向的点能否流到
		for _, d := range dirs {
			if nx, ny := x+d.x, y+d.y; 0 <= nx && nx < m && 0 <= ny && ny < n && heights[nx][ny] >= heights[x][y] {
				dfs(nx, ny, ocean)
			}
		}
	}
	// 左边的m个点 反向搜索
	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
	}
	// 顶部的n-1个点, (0,0) 已搜索
	for j := 1; j < n; j++ {
		dfs(0, j, pacific)
	}
	// 右边的m个点
	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic)
	}
	// 下边的 n-1个点。 (m-1,0) 已搜索
	for j := 0; j < n-1; j++ {
		dfs(m-1, j, atlantic)
	}

	for i, row := range pacific {
		for j, ok := range row {
			if ok && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}
