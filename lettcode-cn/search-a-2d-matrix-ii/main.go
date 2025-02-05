package main

import "fmt"

/**
 * searchMatrix 在二维矩阵中搜索目标值
 * @param matrix 二维矩阵，每行从左到右升序，每列从上到下升序
 * @param target 要搜索的目标值
 * @return bool 是否找到目标值
 * @complexity 时间复杂度：O(m+n)，其中m和n分别是矩阵的行数和列数
 * @complexity 空间复杂度：O(1)
 */
func searchMatrix(matrix [][]int, target int) bool {
	// 从右上角开始搜索
	row := 0
	col := len(matrix[0]) - 1

	for row < len(matrix) && col >= 0 {
		current := matrix[row][col]
		if current == target {
			return true
		}
		if current > target {
			// 当前值太大，排除这一列
			col--
		} else {
			// 当前值太小，排除这一行
			row++
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5
	fmt.Println(searchMatrix(matrix, target))
}
