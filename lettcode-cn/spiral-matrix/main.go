package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var direction = 0

	var visited = make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[0]))
	}

	var res = make([]int, len(matrix)*len(matrix[0]))
	for i, r, c := 0, 0, 0; i < len(res); i++ {
		nextRow, nextCol := r+directions[direction][0], c+directions[direction][1]
		if nextRow < 0 || nextRow >= len(matrix) || nextCol < 0 || nextCol >= len(matrix[0]) || visited[nextRow][nextCol] {
			direction = (direction + 1) % len(directions)
		}
		res[i] = matrix[r][c]
		visited[r][c] = true
		r += directions[direction][0]
		c += directions[direction][1]
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}
