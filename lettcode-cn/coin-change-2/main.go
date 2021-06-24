package main

import "fmt"

//优化前
func change1(amount int, coins []int) int {
	var dp = make([]int ,amount+1, amount+1)
	dp[0] = 1
	for i:=0;i<len(coins);i++ {
		for j:=amount;j>=coins[i];j-- {
			for n:=1;n<= j/coins[i];n++ {
				dp[j] += dp[j-n*coins[i]]
			}

		}
	}
	return dp[amount]
}

func change(amount int, coins []int) int {
	var dp = make([]int ,amount+1, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		//从coin往amount 遍历。就已经算到了使用多个coin的情况。因为使用的个数是从1~n, 所以需要从小往大了遍历
		for j:=coin;j<=amount;j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}

func main()  {
	fmt.Println(change(5,[]int{1,2,5}))
	fmt.Println(change(3,[]int{2}))
	fmt.Println(change(10,[]int{10}))
}