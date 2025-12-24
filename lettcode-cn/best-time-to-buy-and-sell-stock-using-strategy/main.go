package main

func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	prefixSums := make([]int64, n+1)
	baseSums := make([]int64, n+1)
	var base int64
	for i := 0; i < n; i++ {
		prefixSums[i+1] = prefixSums[i] + int64(prices[i])
		base += int64(strategy[i]) * int64(prices[i])
		baseSums[i+1] = baseSums[i] + int64(strategy[i])*int64(prices[i])
	}
	// Adjust profit based on k
	if k > n {
		return base
	}
	var maxProfit int64 = base
	for i := 0; i <= n-k; i++ {
		p1 := i
		p2 := i + k/2 - 1
		p3 := i + k/2
		p4 := i + k - 1
		base1 := (baseSums[p2+1] - baseSums[p1])
		base2 := (baseSums[p4+1] - baseSums[p3])
		p := prefixSums[p4+1] - prefixSums[p3]

		profit := base - base1 - base2 + p
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit

}
