package main

func checkStraightLine(coordinates [][]int) bool {
	for i := 2; i < len(coordinates); i++ {
		if (coordinates[i][1]-coordinates[i-1][1])*(coordinates[i-1][0]-coordinates[i-2][0]) != (coordinates[i-1][1]-coordinates[i-2][1])*(coordinates[i][0]-coordinates[i-1][0]) {
			return false
		}
	}
	return true
}
