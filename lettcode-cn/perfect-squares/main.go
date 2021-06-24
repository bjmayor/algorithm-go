package main

import "math"
// dp 求解
func numSquares1(n int) int {
	coins := []int{}
	for i:=1;i<=n;i++ {
		square := i * i
		if square > n {
			break
		}
		coins = append(coins, square)
	}

	dp := make([]int, n+1, n+1)

	for _, coin := range coins {
		for j:=coin;j<=n;j++ {
			tmp := dp[j-coin]+1
			if dp[j] > 0 {
				dp[j] = min(dp[j], tmp)
			} else {
				dp[j] = tmp
			}
		}
	}
	return dp[n]
}

func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

// 判断是否为完全平方数
func isPerfectSquare(x int) bool {
	y := int(math.Sqrt(float64(x)))
	return y*y == x
}

// 判断是否能表示为 4^k*(8m+7)
func checkAnswer4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}

func numSquares(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}


func main()  {
	println(numSquares(12))
	println(numSquares(13))
}