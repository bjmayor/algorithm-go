package main

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	mod := int(1e9 + 7)

	// dp[i][j][rem]
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}

	// Initialize start point
	dp[0][0][grid[0][0]%k] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for rem := 0; rem < k; rem++ {
				// Current cell value
				val := grid[i][j]
				// Calculate new remainder
				newRem := (rem + val) % k

				// From top
				if i > 0 {
					dp[i][j][newRem] = (dp[i][j][newRem] + dp[i-1][j][rem]) % mod
				}
				// From left
				if j > 0 {
					dp[i][j][newRem] = (dp[i][j][newRem] + dp[i][j-1][rem]) % mod
				}
			}
		}
	}

	return dp[m-1][n-1][0]
}
