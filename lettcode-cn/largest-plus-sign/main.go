package main

func orderOfLargestPlusSign(n int, mines [][]int) int {
	var arm int
	var m map[[2]int]int = make(map[[2]int]int)
	for _, v := range mines {
		m[[2]int{v[0], v[1]}] = 0
	}
	var equal1 = func(x, y int) bool {
		if _, ok := m[[2]int{x, y}]; ok {
			return false
		}
		return true
	}
	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			if !equal1(i, j) {
				continue
			}
			l := 0
			for k := 0; k < i; k++ {
				if equal1(k, j) {
					l++
				} else {
					break
				}
			}
			if l < arm {
				continue
			}
			r := 0
			for k := i + 1; k < n; k++ {
				if equal1(k, j) {
					r++
				} else {
					break
				}
			}
			if r < arm {
				continue
			}

			x := 0
			for k := 0; k < j; k++ {
				if equal1(i, k) {
					x++
				} else {
					break
				}
			}
			if x < arm {
				continue
			}

			y := 0
			for k := j + 1; k < n; k++ {
				if equal1(i, k) {
					y++
				} else {
					break
				}
			}
			if y < arm {
				continue
			}
			v := min(l, r)
			v = min(v, x)
			v = min(v, y)
			if v > arm {
				arm = v
			}
		}
	}

	return arm + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(orderOfLargestPlusSign(2, [][]int{{0, 0}, {0, 1}, {1, 0}}))
}
