package main

func countCollisions(directions string) int {
	total := 0
	rCount := 0
	hasCar := false
	
	for _, char := range directions {
		switch char {
		case 'R':
			hasCar = true
			rCount++
		case 'L':
			if hasCar {
				total += rCount + 1
				rCount = 0
			}
		case 'S':
			hasCar = true
			total += rCount
			rCount = 0
		}
	}
	return total
}
