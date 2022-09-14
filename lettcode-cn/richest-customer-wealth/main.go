package main

func maximumWealth(accounts [][]int) int {
	var ans int
	for _, row := range accounts {
		var s int
		for _, v := range row {
			s += v
		}
		if s > ans {
			ans = s
		}
	}
	return ans
}
