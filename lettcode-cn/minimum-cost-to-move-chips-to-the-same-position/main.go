package main

import "math"

func minCostToMoveChips(position []int) int {
	return greed(position)
}
func baoli(position []int) int {
	m := math.MaxInt32
	for _, p := range position {
		n := 0
		for _, q := range position {
			n += cost(p, q)
		}
		if n < m {
			m = n
		}
	}
	return m
}

func baoli2(position []int) int {
	d := map[int]int{} // key为位置，value为多少个
	for _, p := range position {
		d[p]++
	}
	m := math.MaxInt32
	for p, num := range d {
		n := 0
		for q, _ := range d {
			n += cost(p, q) * num
		}
		if n < m {
			m = n
		}
	}
	return m
}

func baoli3(position []int) int {
	d := map[int]int{} // key为位置，value为多少个
	for _, p := range position {
		d[p]++
	}
	m := math.MaxInt32
	for _, q := range []int{1, 2} {
		n := 0
		for p, num := range d {

			n += cost(p, q) * num

		}
		if n < m {
			m = n
		}
	}

	return m
}

func greed(position []int) int {
	costs := []int{0, 0}
	for _, p := range position {
		costs[p%2]++
	}
	if costs[0] > costs[1] {
		return costs[1]
	}
	return costs[0]
}
func cost(p, q int) int {
	v := (p - q) % 2
	if v < 0 {
		return -v
	}
	return v
}

func main() {

	println(minCostToMoveChips([]int{10, 3, 3, 1, 6, 2, 1, 10, 6, 6})) // 4
	println(minCostToMoveChips([]int{1, 2, 3}))                        // 1
	println(minCostToMoveChips([]int{2, 2, 2, 3, 3}))                  // 2
	println(minCostToMoveChips([]int{1, 1, 1000000000}))               //1

}
