package main

func finalPrices2(prices []int) []int {
	result := make([]int, len(prices))
	for i, price := range prices {
		m := 0
		for _, q := range prices[i+1:] {
			if q <= price {
				m = q
				break
			}
		}
		result[i] = price - m
	}
	return result
}

func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	st := []int{}
	for i := len(prices) - 1; i >= 0; i-- {
		for len(st) > 0 && prices[i] < st[len(st)-1] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = prices[i] - st[len(st)-1]
		} else {
			ans[i] = prices[i]
		}
		st = append(st, prices[i])
	}
	return ans
}