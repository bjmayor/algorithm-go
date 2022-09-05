package main

func numSpecial(mat [][]int) int {
	rowsSum := make([]int, len(mat))
	colsSum := make([]int, len(mat[0]))
	for i, row := range mat {
		for j, x := range row {
			rowsSum[i] += x
			colsSum[j] += x
		}
	}

	var ans int
	for i, row := range mat {
		for j, v := range row {
			if v == 1 && rowsSum[i] == 1 && colsSum[j] == 1 {
				ans++
			}
		}
	}
	return ans
}

func numSpecial2(mat [][]int) (ans int) {
	for i, row := range mat {
		cnt1 := 0
		for _, x := range row {
			cnt1 += x
		}
		if i == 0 { // 第一行只有 和>1时才需要 累加下。
			cnt1--
		}
		if cnt1 > 0 {
			for j, x := range row {
				if x == 1 {
					mat[0][j] += cnt1
				}
			}
		}
	}
	for _, x := range mat[0] {
		if x == 1 {
			ans++
		}
	}
	return
}

func main() {
	println(numSpecial2([][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}))
}
