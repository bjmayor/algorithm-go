package main

import "fmt"

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	// 定义[i][j]
	// i 表示前i块瓷砖
	// j 表示用了 j块地毯
	dp := make([][]int, len(floor)+1)
	dp[0] = make([]int, numCarpets+1)
	// 初始化
	for i := 0; i <= numCarpets; i++ {
		dp[0][i] = 0
	}
	// 转移
	for i := 1; i <= len(floor); i++ {
		dp[i] = make([]int, numCarpets+1)
		num := 0
		if floor[i-1] == '1' {
			num += 1
		}
		dp[i][0] = dp[i-1][0] + num
		for j := 1; j <= numCarpets; j++ {
			dp[i][j] = dp[i-1][j] + num
			// 选择j 块地摊时，有几种和情况
			//1. j 盖住地i块。
			//2. j 没有盖住i块。

			if i <= carpetLen {
				// 就这 1 块，我都盖住了
				dp[i][j] = 0
			} else {
				dp[i][j] = min(dp[i][j], dp[i-carpetLen][j-1])
			}

		}

	}

	return dp[len(floor)][numCarpets]
}

func main() {
	floor := "10110101"
	numCarpets := 2
	carpetLen := 2
	fmt.Println(minimumWhiteTiles(floor, numCarpets, carpetLen))

}
