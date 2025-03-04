package main

import "fmt"

func checkPartitioning(s string) bool {
	n := len(s)
	// 预处理所有子串是否为回文
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true // 单个字符是回文
	}

	// 计算所有子串的回文状态
	for span := 2; span <= n; span++ {
		for i := 0; i <= n-span; i++ {
			j := i + span - 1
			if span == 2 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
		}
	}

	// 使用预处理结果检查三段式回文
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if dp[0][i-1] && dp[i][j-1] && dp[j][n-1] {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(checkPartitioning("abcbdd"))  //true
	fmt.Println(checkPartitioning("bcbddxy")) //false
	fmt.Println(checkPartitioning("acab"))    //false
}
