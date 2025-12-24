package main

func minDeletionSize(strs []string) int {
	n := len(strs)
	m := len(strs[0])
	dp := make([]int, m)
	dp[0] = 1
	maxLen := 1
	var isValid = func(i, j int) bool {
		for k := 0; k < n; k++ {
			if strs[k][i] > strs[k][j] {
				return false
			}
		}
		return true
	}
	for i := 1; i < m; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if isValid(j, i) {
				dp[i] = max(dp[i], dp[j]+1)
				if dp[i] > maxLen {
					maxLen = dp[i]
				}
			}
		}
	}
	return m - maxLen
}
