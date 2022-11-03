package main

func orangesRotting(grid [][]int) int {
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(i, j int, idx int)
	dfs = func(i, j int, idx int) {
		for _, dir := range dirs {
			row, col := i+dir[0], j+dir[1]
			if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) {
				if grid[row][col] > 0 {
					if grid[row][col] == 1 {
						grid[row][col] = idx
						dfs(row, col, idx+1)
					} else {
						if grid[row][col] > idx {
							grid[row][col] = idx
							dfs(row, col, idx+1)
						}
					}
				}
			}
		}
	}
	for i, rows := range grid {
		for j, v := range rows {
			if v == 2 {
				dfs(i, j, v+1)
			}
		}
	}
	var ans = 0
	for _, rows := range grid {
		for _, v := range rows {
			if v == 1 {
				return -1
			} else if v > ans {
				ans = v
			}
		}
	}
	if ans > 2 {
		return ans - 2
	}
	return 0
}

func main() {
	//println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})) // 4,
	//println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}})) // -1
	//println(orangesRotting([][]int{{0, 2}}))                          // 0
	println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}})) // 2,
}
