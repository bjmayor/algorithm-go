package main

import "fmt"

func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 || k<1{
		return 0
	}
	if k > len(prices) {
		return maxProfitNoLimit(prices)
	}
	result := make([]int, 2*k, 2*k)
	for i:=0;i<2*k;i+=2 {
		result[i] = -prices[0]
		result[i+1] = 0
	}

	for i:=1;i<len(prices);i++ {
		for j:=0;j<2*k;j+=2{
			leftProfit := 0
			if j > 0 {
				leftProfit = result[j-1]
			}
			result[j] = max(result[j], leftProfit-prices[i])
			result[j+1] = max(result[j+1], result[j]+prices[i])
		}
	}

	return result[2*k-1]
}

func maxProfitNoLimit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	for i := 1;i<len(prices); i++ {
		if prices[i]>prices[i-1]{
			profit += prices[i]-prices[i-1]
		}
	}
	return profit
}

func max(x,y int) int  {
	if x > y {
		return x
	}
	return y
}

func main()  {
	tests := []struct{
		Nums []int
		k int
		Expected int
	} {
		//{ []int{3,3,5,0,0,3,1,4},6},
		//{ []int{1,2,3,4,5},4},
		{ []int{2,4,1},2,2},
		{ []int{3,2,6,5,0,3},2,7},
	}
	for _, t := range tests {
		var real = maxProfit(t.k, t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
