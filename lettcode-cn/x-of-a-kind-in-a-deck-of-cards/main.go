package main

func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)
	for _, v := range deck {
		m[v]++
	}
	var nums []int
	for _, v := range m {
		nums = append(nums, v)
	}
	g := -1
	for i, v := range nums {
		if i == 0 {
			g = v
		} else {
			g = gcd(g, v)
		}
	}
	return g >= 2
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}
