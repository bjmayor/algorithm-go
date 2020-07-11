package main

import "fmt"

func isPathCrossing(path string) bool {
	p := make(map[int]bool)
	x:=0
	y:=0
	p[0] = true
	for _, v := range path {
		switch v {
		case 'N':
			y++
		case 'S':
			y--
		case 'W':
			x--
		case 'E':
			x++
		}
		hash := 10000 * x +y
		if p[hash] {
			return true
		}
		p[hash] = true
	}
	return false
}

func main()  {
	fmt.Println(isPathCrossing("NESWW")) //true
}


