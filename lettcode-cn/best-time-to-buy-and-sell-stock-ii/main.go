package main

import "fmt"

//贪心算法
func maxProfit(prices []int) int {
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

func main()  {
	tests := []struct{
		Nums []int
		Expected int
	} {
		{ []int{7,1,5,3,6,4},7},
		{ []int{1,2,3,4,5},4},
	}
	for _, t := range tests {
		var real = maxProfit(t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}

