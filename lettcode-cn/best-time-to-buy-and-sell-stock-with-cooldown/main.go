package main

import "math"

func maxProfit(prices []int) int {
	n := len(prices)
	dp_i_0 := 0 //空仓的利润
	dp_i_1 := math.MinInt32 //持仓的利润
	dp_pre_0 := 0; // 代表 dp[i-2][0] //上次卖出是的利润
	for i := 0; i < n; i++ {
		temp := dp_i_0;
		dp_i_0 = max(dp_i_0, dp_i_1 + prices[i]); //第i天空仓，可能是因为之前就空仓，也可能是因为第i天卖了。
		dp_i_1 = max(dp_i_1, dp_pre_0 - prices[i]);//第i天满仓，可能是之前就满仓，也可能是第i-2天空仓，然后第i天买了
		dp_pre_0 = temp;
	}
	return dp_i_0;
}

func max(x,y int)  int {
	if x > y {
		return x
	}
	return y
}


