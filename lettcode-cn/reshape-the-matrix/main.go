package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	var ans [][]int = make([][]int, r)
	for i := 0; i < r; i++ {
		ans[i] = make([]int, c)
		for j := 0; j < c; j++ {
			num := i*c + j
			ans[i][j] = mat[num/n][num%n]
		}
	}
	return ans
}
