package main

import "fmt"

/*
2711. 对角线上不同值的数量差

给你一个下标从 0 开始、大小为 m x n 的二维矩阵 grid ，请你求解大小同样为 m x n 的答案矩阵 answer 。

矩阵 answer 中每个单元格 (r, c) 的值可以按下述方式进行计算：

令 topLeft[r][c] 为矩阵 grid 中单元格 (r, c) 左上角对角线上 不同值 的数量。
令 bottomRight[r][c] 为矩阵 grid 中单元格 (r, c) 右下角对角线上 不同值 的数量。
然后 answer[r][c] = |topLeft[r][c] - bottomRight[r][c]| 。

返回矩阵 answer 。

矩阵对角线 是从最顶行或最左列的某个单元格开始，向右下方向走到矩阵末尾的对角线。

如果单元格 (r1, c1) 和单元格 (r, c) 属于同一条对角线且 r1 < r ，则单元格 (r1, c1) 属于单元格 (r, c) 的左上对角线。类似地，可以定义右下对角线。
*/
func differenceOfDistinctValues(grid [][]int) [][]int {
	rows := len(grid)
	cols := len(grid[0])
	answer := make([][]int, rows)
	for i := range answer {
		answer[i] = make([]int, cols)
	}

	// 处理每条对角线
	// 从第一行开始的对角线
	for c := 0; c < cols; c++ {
		processDiagonal(grid, 0, c, answer)
	}
	// 从第一列开始的对角线（除了已经处理过的(0,0)）
	for r := 1; r < rows; r++ {
		processDiagonal(grid, r, 0, answer)
	}

	return answer
}

// 处理单条对角线
func processDiagonal(grid [][]int, startR, startC int, answer [][]int) {
	rows, cols := len(grid), len(grid[0])
	// 收集对角线上的所有坐标
	var positions [][2]int
	r, c := startR, startC
	for r < rows && c < cols {
		positions = append(positions, [2]int{r, c})
		r++
		c++
	}

	n := len(positions)
	if n == 0 {
		return
	}

	// 从左到右计算每个位置左上方的不同值
	topLeft := make([]int, n)
	seen := make(map[int]bool)
	for i := 0; i < n; i++ {
		if i > 0 {
			topLeft[i] = len(seen)
		}
		r, c := positions[i][0], positions[i][1]
		seen[grid[r][c]] = true
	}

	// 从右到左计算每个位置右下方的不同值
	bottomRight := make([]int, n)
	seen = make(map[int]bool)
	for i := n - 1; i >= 0; i-- {
		if i < n-1 {
			bottomRight[i] = len(seen)
		}
		r, c := positions[i][0], positions[i][1]
		seen[grid[r][c]] = true
	}

	// 填充答案
	for i := 0; i < n; i++ {
		r, c := positions[i][0], positions[i][1]
		diff := topLeft[i] - bottomRight[i]
		if diff < 0 {
			diff = -diff
		}
		answer[r][c] = diff
	}
}

func main() {
	grid := [][]int{{1, 2, 3}, {3, 1, 5}, {3, 2, 1}}
	fmt.Println(differenceOfDistinctValues(grid))
}
