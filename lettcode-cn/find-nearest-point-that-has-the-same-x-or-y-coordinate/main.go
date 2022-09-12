package main

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	var v = math.MaxInt64
	var ans = -1
	for i, p := range points {
		if p[0] == x || p[1] == y {
			dis := abs(p[1]-y) + abs(p[0]-x)
			if dis < v {
				v = dis
				ans = i
			}
		}
	}
	return ans
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
