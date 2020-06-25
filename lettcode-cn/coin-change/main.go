package main

import (
	"fmt"
)
//dp求解
func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i:=1;i<=amount;i++ {
		dp[i] = amount+1
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i>=coin {
				tmp := dp[i-coin]+1
				if tmp < dp[i] {
					dp[i] = tmp
				}
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}




func main() {
	fmt.Println(coinChange([]int{1, 2, 5,}, 6))  //2
	fmt.Println(coinChange([]int{1, 2, 5,}, 11)) //3
	fmt.Println(coinChange([]int{2}, 3))         //-1
	fmt.Println(coinChange([]int{1}, 0))         //0
	fmt.Println(coinChange([]int{2}, 4))         //2
	fmt.Println(coinChange([]int{1,2,5}, 100))         //20
}
