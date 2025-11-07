package main

import "fmt"

/*
*
给定一个字符串数组 words ，和一个数组 groups ，两个数组长度都是 n 。

两个长度相等字符串的 汉明距离 定义为对应位置字符 不同 的数目。

你需要从下标 [0, 1, ..., n - 1] 中选出一个 最长子序列 ，将这个子序列记作长度为 k 的 [i0, i1, ..., ik - 1] ，它需要满足以下条件：

相邻 下标对应的 groups 值 不同。即，对于所有满足 0 < j + 1 < k 的 j 都有 groups[ij] != groups[ij + 1] 。
对于所有 0 < j + 1 < k 的下标 j ，都满足 words[ij] 和 words[ij + 1] 的长度 相等 ，且两个字符串之间的 汉明距离 为 1 。
请你返回一个字符串数组，它是下标子序列 依次 对应 words 数组中的字符串连接形成的字符串数组。如果有多个答案，返回任意一个。

子序列 指的是从原数组中删掉一些（也可能一个也不删掉）元素，剩余元素不改变相对位置得到的新的数组。

注意：words 中的字符串长度可能 不相等 。
*/
/**
 * 获取满足条件的最长子序列对应的字符串数组。
 * 条件：
 * 1. 相邻下标对应 groups 不同。
 * 2. 相邻字符串长度相等且汉明距离为 1。
 * 3. 返回任意一个最长子序列。
 * @param words 字符串数组
 * @param groups 分组数组
 * @return 最长合法子序列对应的字符串数组
 */
func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	// dp[i]: 以i结尾的最长子序列长度
	dp := make([]int, n)
	// prev[i]: 以i结尾的最长子序列的前驱下标
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		prev[i] = -1
	}

	// 汉明距离为1的判断
	hamming1 := func(a, b string) bool {
		if len(a) != len(b) {
			return false
		}
		cnt := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				cnt++
				if cnt > 1 {
					return false
				}
			}
		}
		return cnt == 1
	}

	// 动态规划
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if groups[j] != groups[i] && hamming1(words[j], words[i]) {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					prev[i] = j
				}
			}
		}
	}

	// 找到最长子序列的结尾下标
	maxLen := 0
	end := 0
	for i := 0; i < n; i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
			end = i
		}
	}

	// 回溯得到路径
	res := make([]string, 0, maxLen)
	for end != -1 {
		res = append(res, words[end])
		end = prev[end]
	}
	// 反转
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func main() {
	fmt.Println(getWordsInLongestSubsequence([]string{"a", "b", "c", "d"}, []int{1, 2, 3, 4}))
	fmt.Println(getWordsInLongestSubsequence([]string{"bab", "dab", "cab"}, []int{1, 2, 2}))
}
