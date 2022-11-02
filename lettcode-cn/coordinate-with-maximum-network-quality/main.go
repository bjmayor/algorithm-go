package main

import (
	"fmt"
	"math"
)

func bestCoordinate(towers [][]int, radius int) []int {
	var xMax, yMax, cx, cy, maxQuality int
	for _, t := range towers {
		xMax = max(xMax, t[0])
		yMax = max(yMax, t[1])
	}
	for x := 0; x <= xMax; x++ {
		for y := 0; y <= yMax; y++ {
			quality := 0
			for _, t := range towers {
				d := (x-t[0])*(x-t[0]) + (y-t[1])*(y-t[1])
				if d <= radius*radius {
					quality += int(float64(t[2]) / (1 + math.Sqrt(float64(d))))
				}
			}
			if quality > maxQuality {
				cx, cy, maxQuality = x, y, quality
			}
		}
	}
	return []int{cx, cy}
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	fmt.Println(bestCoordinate([][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}, 2))
	fmt.Println(bestCoordinate([][]int{{23, 11, 21}}, 9))
	fmt.Println(bestCoordinate([][]int{{1, 2, 13}, {2, 1, 7}, {0, 1, 9}}, 2))
}
