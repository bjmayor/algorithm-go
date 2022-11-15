package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) (ans int) {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	for i := 0; i < len(boxTypes); i++ {
		used := min(boxTypes[i][0], truckSize)
		if used > 0 {
			ans += used * boxTypes[i][1]
			truckSize -= used
			if truckSize == 0 {
				return
			}
		}
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
}