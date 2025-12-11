package main

import "fmt"

func countCoveredBuildings(n int, buildings [][]int) int {
	xMap := make(map[int][2]int)
	yMap := make(map[int][2]int)
	count := 0

	// 初始化映射
	for _, building := range buildings {
		x, y := building[0], building[1]
		if _, exists := xMap[x]; !exists {
			xMap[x] = [2]int{y, y}
		} else {
			xMap[x] = [2]int{min(xMap[x][0], y), max(xMap[x][1], y)}
		}
		if _, exists := yMap[y]; !exists {
			yMap[y] = [2]int{x, x}
		} else {
			yMap[y] = [2]int{min(yMap[y][0], x), max(yMap[y][1], x)}
		}
	}

	for _, building := range buildings {
		if xMap[building[0]][0] < building[1] && xMap[building[0]][1] > building[1] {
			if yMap[building[1]][0] < building[0] && yMap[building[1]][1] > building[0] {
				count++
			}
		}
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	buildings := [][]int{
		{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4},
		{0, 2}, {1, 2}, {3, 2}, {4, 2},
	}
	result := countCoveredBuildings(5, buildings)
	fmt.Printf("被覆盖的建筑数量: %d\n", result)
}
