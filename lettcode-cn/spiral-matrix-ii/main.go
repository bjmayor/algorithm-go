package main

import "fmt"

func generateMatrix(n int) [][]int {
	var res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	left, right, top, botton, idx := 0, n-1, 0, n-1, 1
	for left <= right && top <= botton {
		// 顶部 左到右
		for c := left; c <= right; c++ {
			res[top][c] = idx
			idx++
		}
		// 右边 上到下
		for r := top + 1; r <= botton; r++ {
			res[r][right] = idx
			idx++
		}
		// 底部 右到左
		for c := right - 1; c >= left; c-- {
			res[botton][c] = idx
			idx++
		}
		// 左边 下到上
		for r := botton - 1; r >= top+1; r-- {
			res[r][left] = idx
			idx++
		}
		left++
		right--
		top++
		botton--
	}
	return res
}

func main() {
	fmt.Println(generateMatrix(3))
}
