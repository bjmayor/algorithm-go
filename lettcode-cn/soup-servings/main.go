package main

import "fmt"

func soupServings(n int) float64 {
	if n >= 5000 {
		return 1
	}
	var mem = map[string][2]float64{}
	var diliver func(a, b int) (first float64, same float64)
	diliver = func(a, b int) (first float64, same float64) {
		key := fmt.Sprintf("%d %d", a, b)
		defer func() {
			mem[key] = [2]float64{first, same}
		}()
		if v, ok := mem[key]; ok {
			return v[0], v[1]
		}
		if a == 0 {
			if b == 0 {
				same = 1
				return
			}
			first = 1
			return
		}
		if b == 0 {
			return
		}
		todo := [][]int{{100, 0}, {75, 25}, {50, 50}, {25, 75}}
		for _, item := range todo {
			p1, p2 := diliver(max(a-item[0], 0), max(b-item[1], 0))
			first += p1
			same += p2
		}
		first *= 0.25
		same *= 0.25
		return
	}
	first, same := diliver(n, n)
	return first + same/2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(soupServings(50))
	println(soupServings(100))
	println(soupServings(800))
}
