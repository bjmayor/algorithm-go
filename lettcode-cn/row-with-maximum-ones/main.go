package main

import "fmt"

func rowAndMaximumOnes(mat [][]int) []int {
	var ans []int = []int{0, 0}
	for i, row := range mat {
		sum := 0
		for _, v := range row {
			sum += v
		}
		if sum > ans[1] {
			ans = []int{i, sum}
		}
	}
	return ans
}

func main() {
	mat := [][]int{{0, 1}, {1, 0}}
	fmt.Println(rowAndMaximumOnes(mat))
}
