package main


import "fmt"

//单调栈 算法
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	low := prices[0]
	profile1 := 0
	leftAferBuy2 := -low
	profile2 := 0
	for i := 1; i< len(prices); i++ {
		low = min(low, prices[i]) //左边最小价格
		profile1 = max(profile1, prices[i]-low) //第一天卖能产生的最大利润。
		leftAferBuy2 = max(leftAferBuy2, profile1-prices[i]) //这里其实是求第二次交易的 左边最小价格
		profile2 = max(profile2, leftAferBuy2+prices[i])
	}
	return profile2
}

func min(x,y int) int  {
	if x < y {
		return x
	}
	return y
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
		Expected int
	} {
		//{ []int{3,3,5,0,0,3,1,4},6},
		//{ []int{1,2,3,4,5},4},
		{ []int{2,1,2,0,1},2},
	}
	for _, t := range tests {
		var real = maxProfit(t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}

