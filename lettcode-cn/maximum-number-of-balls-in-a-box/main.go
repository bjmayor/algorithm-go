package main

func countBalls(lowLimit int, highLimit int) (ans int) {
	m := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		idx := countNum(i)
		m[idx]++
		if m[idx] > ans {
			ans = m[idx]
		}
	}
	return
}

func countNum(i int) (v int) {
	for i > 0 {
		v += i % 10
		i /= 10
	}
	return
}
