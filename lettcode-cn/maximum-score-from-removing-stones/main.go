package main

func maximumScore(a, b, c int) int {
	sum := a + b + c
	maxVal := max(max(a, b), c)
	if sum < maxVal*2 { // a+b+c < c+c
		return sum - maxVal
	} else {
		return sum / 2
	}
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	println(maximumScore(2, 4, 6)) // 6
	println(maximumScore(4, 4, 6)) // 7
	println(maximumScore(1, 8, 8)) // 8
}
