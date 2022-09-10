package main

import "fmt"

func generate(numRows int) [][]int {
	var ans = [][]int{{1}}
	for i := 2; i <= numRows; i++ {
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		for j := 1; j <= i-2; j++ {
			row[j] = ans[i-2][j-1] + ans[i-2][j]
		}
		ans = append(ans, row)
	}
	return ans
}

func main() {
	fmt.Println(generate(5))
}
