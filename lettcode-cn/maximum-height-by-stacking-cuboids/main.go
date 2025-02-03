package main

import "sort"

func maxHeight(cuboids [][]int) (ans int) {
	for i, cuboid := range cuboids {
		sort.Ints(cuboid)
		cuboids[i] = cuboid
	}
	sort.Slice(cuboids, func(i, j int) bool {
		return cuboids[i][0]+cuboids[i][1]+cuboids[i][2] < cuboids[j][0]+cuboids[j][1]+cuboids[j][2]
	})
	dp := make([]int, len(cuboids))
	for i := 0; i < len(cuboids); i++ {
		dp[i] = cuboids[i][2]
		for j := 0; j < i; j++ {
			if cuboids[j][0] <= cuboids[i][0] && cuboids[j][1] <= cuboids[i][1] && cuboids[j][2] <= cuboids[i][2] {
				dp[i] = max(dp[i], dp[j]+cuboids[i][2])
			}
		}
		ans = max(ans, dp[i])
	}
	return
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
