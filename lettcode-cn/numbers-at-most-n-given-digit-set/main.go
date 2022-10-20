package main

import "strconv"

func atMostNGivenDigitSet(digits []string, n int) int {
	nstr := strconv.Itoa(n)
	k := len(nstr)
	dp := make([][2]int, k+1) // dp[i][0] < nstr[0:i] dp[i][1] = nstr[0:i]
	dp[0][1] = 1
	for i := 0; i < k; i++ {
		for _, v := range digits {
			if v[0] == nstr[i] {
				dp[i+1][1] = dp[i][1]
			} else if v[0] < nstr[i] {
				dp[i+1][0] += dp[i][1]
			} else {
				break
			}
		}
		if i > 0 {
			dp[i+1][0] += len(digits) + dp[i][0]*len(digits)
		}
	}
	return dp[k][0] + dp[k][1]
}

func main() {
	println(atMostNGivenDigitSet([]string{"7"}, 8))                  //1
	println(atMostNGivenDigitSet([]string{"2", "5", "8", "9"}, 6))   //2
	println(atMostNGivenDigitSet([]string{"5", "7", "8"}, 59))       //6
	println(atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100)) //20
	println(atMostNGivenDigitSet([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, 74546987))
}
