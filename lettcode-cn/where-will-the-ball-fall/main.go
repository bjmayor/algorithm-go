package main

import "fmt"

func findBall(grid [][]int) []int {
	var ans []int = make([]int, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		row := 0
		col := i
		for row < len(grid) {
			switch grid[row][col] {
			case 1:
				if col == len(grid[0])-1 {
					ans[i] = -1
				} else if grid[row][col+1] == -1 {
					ans[i] = -1
				} else {
					row += 1
					col += 1
				}
			case -1:
				if col == 0 {
					ans[i] = -1
				} else if grid[row][col-1] == 1 {
					ans[i] = -1
				} else {
					row += 1
					col -= 1
				}
			}
			if ans[i] == -1 {
				break
			}
		}
		if ans[i] == 0 {
			ans[i] = col
		}
	}
	return ans
}

func main() {
	//fmt.Println(findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}))

	fmt.Println(findBall([][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}))
}
