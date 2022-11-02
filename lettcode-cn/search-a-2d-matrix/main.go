package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) >> 1
		row := mid / n
		col := mid % n
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			right = row*n + col - 1
		} else {
			left = row*n + col + 1
		}
	}
	return false
}

func main() {
	println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 3))
}
