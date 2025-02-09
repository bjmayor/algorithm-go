package main

import "fmt"

/**
https://leetcode.cn/problems/partition-labels/?envType=study-plan-v2&envId=top-100-liked

763. 划分字母区间
给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。

注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。

返回一个表示每个字符串片段的长度的列表。


*/

/*
* partitionLabels 划分字母区间
* 算法思想:
* 1. 记录每个字母的最后出现位置
* 2. 遍历字符串, 使用last变量跟踪当前片段的结束位置
* 3. 当到达当前片段的结束位置时, 将片段长度添加到结果列表
*
* 时间复杂度: O(n) - 遍历字符串两次
* 空间复杂度: O(1) - 只使用常数空间
 */
func partitionLabels(s string) []int {
	// 记录每个字母的最后出现位置
	last := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		last[s[i]] = i
	}

	var result []int
	start, end := 0, 0

	// 遍历字符串
	for i := 0; i < len(s); i++ {
		// 更新当前片段的结束位置
		end = max(end, last[s[i]])
		// 如果到达当前片段的结束位置
		if i == end {
			// 计算片段长度并添加到结果
			result = append(result, end-start+1)
			start = end + 1 // 更新下一个片段的起始位置
		}
	}

	return result
}

// 辅助函数: 返回较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij")) // [9 7 8]
}
