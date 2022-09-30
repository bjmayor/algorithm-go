package main

func setZeroes(matrix [][]int) {
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[0]))
	}
	var dfs func(x, y int)
	dfs = func(i, j int) {
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		for p := 0; p < len(matrix); p++ {
			if matrix[p][j] == 0 {
				dfs(p, j)
			}
			matrix[p][j] = 0
			visited[p][j] = true
		}
		for q := 0; q < len(matrix[0]); q++ {
			if matrix[i][q] == 0 {
				dfs(i, q)
			}
			matrix[i][q] = 0
			visited[i][q] = true
		}
	}
	for i, rows := range matrix {
		for j, col := range rows {

			if col == 0 {
				dfs(i, j)
			}
		}
	}
}
