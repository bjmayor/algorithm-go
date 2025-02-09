package main

import "fmt"

/**
 * @description: 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵。
 * @param {int} n 矩阵大小
 * @return {[][]int} 生成的螺旋矩阵
 */
func generateMatrix(n int) [][]int {
	// 初始化 n x n 的二维数组
	var res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	// 定义四个边界和当前要填入的数字
	left, right, top, botton, idx := 0, n-1, 0, n-1, 1

	// 当左边界小于等于右边界且上边界小于等于下边界时继续循环
	for left <= right && top <= botton {
		// 顶部从左到右填充
		for c := left; c <= right; c++ {
			res[top][c] = idx
			idx++
		}
		// 右边从上到下填充
		for r := top + 1; r <= botton; r++ {
			res[r][right] = idx
			idx++
		}
		// 底部从右到左填充
		for c := right - 1; c >= left; c-- {
			res[botton][c] = idx
			idx++
		}
		// 左边从下到上填充
		for r := botton - 1; r >= top+1; r-- {
			res[r][left] = idx
			idx++
		}
		// 缩小边界范围,进入内层
		left++
		right--
		top++
		botton--
	}
	return res
}

/**
 * @description: 主函数,用于测试
 */
func main() {
	fmt.Println(generateMatrix(3))
}
