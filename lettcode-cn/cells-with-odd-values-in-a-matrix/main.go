package main
func oddCells(m int, n int, indices [][]int) int {
	matrix := make([][]int,0)
	for i:=0;i<m;i++ {
		matrix = append(matrix, make([]int, n,n))
	}
	for r:=0;r<len(indices);r++ {
		row := indices[r][0]
		col := indices[r][1]
		for c:=0;c<n;c++ {
			matrix[row][c] +=1
		}
		for j:=0;j<m;j++ {
			matrix[j][col] +=1
		}
	}
	odd := 0
	for r:=0;r<m;r++ {
		for c:=0;c<n;c++ {
			if matrix[r][c]	%2 == 1 {
				odd++
			}
		}
	}
	return odd
}

func oddCells2(m int, n int, indices [][]int) int {
	rows := make([]bool,m,m)
	cols := make([]bool,n,n)
	for r:=0;r<len(indices);r++ {
		row := indices[r][0]
		col := indices[r][1]
		rows[row]=!rows[row]
		cols[col]=!cols[col]
	}
	oddRows := 0
	oddCols:= 0
	for _, odd := range rows {
		if odd {
			oddRows+=1
		}
	}
	for _, odd := range cols {
		if odd {
			oddCols+=1
		}
	}
	return oddRows*n+oddCols*m-2*oddRows*oddCols
}

func main()  {
	println(oddCells2(2,3,[][]int{{0,1},{1,1}}))
}
