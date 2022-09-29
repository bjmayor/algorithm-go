package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	dirs := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := map[[2]int]bool{}
	target := image[sr][sc]
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= len(image) || y < 0 || y >= len(image[0]) {
			return
		}
		if !visited[[2]int{x, y}] {
			visited[[2]int{x, y}] = true
			if image[x][y] == target {
				image[x][y] = color
				for _, dir := range dirs {
					nx, ny := x+dir[0], y+dir[1]
					dfs(nx, ny)
				}
			}
		}
	}
	dfs(sr, sc)
	return image
}

func main() {
	fmt.Println(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
}
