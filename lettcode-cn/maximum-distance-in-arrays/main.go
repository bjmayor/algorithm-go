package main

import "fmt"

func maxDistance(arrays [][]int) int {
	m := len(arrays)
	res := 0
	left, right := arrays[0][0], arrays[0][len(arrays[0])-1]
	for i := 1; i < m; i++ {
		temp1 := abs(right, arrays[i][0])
		temp2 := abs(left, arrays[i][len(arrays[i])-1])
		res = max(res, temp1)
		res = max(res, temp2)
		left = min(left, arrays[i][0])
		right = max(right, arrays[i][len(arrays[i])-1])

	}
	return res
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func main() {
	// fmt.Println(maxDistance([][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}})) //4
	// fmt.Println(maxDistance([][]int{{1, 4}, {0, 5}}))               //4
	// fmt.Println(maxDistance([][]int{{1, 5}, {3, 4}}))               //3
	// [[-1,1],[-3,1,4],[-2,-1,0,2]]
	fmt.Println(maxDistance([][]int{{-1, 1}, {-3, 1, 4}, {-2, -1, 0, 2}})) //6
}
