package main

import "fmt"

/*
https://leetcode.cn/problems/longest-common-subsequence/description/?envType=study-plan-v2&envId=top-100-liked
1143. 最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
*/

/*
* longestCommonSubsequence 计算两个字符串的最长公共子序列长度
* 算法思想:
* 1. 使用动态规划, dp[i][j] 表示 text1 的前 i 个字符和 text2 的前 j 个字符的 LCS 长度
* 2. 状态转移:
*    - 如果 text1[i-1] == text2[j-1], dp[i][j] = dp[i-1][j-1] + 1
*    - 否则, dp[i][j] = max(dp[i-1][j], dp[i][j-1])
* 3. 初始化:
*    - dp[0][j] 和 dp[i][0] 都为 0
*
* 时间复杂度: O(m*n) - m 和 n 分别为 text1 和 text2 的长度
* 空间复杂度: O(m*n) - dp 数组
 */
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// 创建 dp 数组
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 填充 dp 数组
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// 辅助函数: 返回较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2)) // 输出: 3
}
