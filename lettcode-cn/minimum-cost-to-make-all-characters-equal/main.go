package main

import "fmt"

/*
2712. 使所有字符相等的最小成本

给你一个下标从 0 开始、长度为 n 的二进制字符串 s ，你可以对其执行两种操作：

选中一个下标 i 并且反转从下标 0 到下标 i（包括下标 0 和下标 i ）的所有字符，成本为 i + 1 。
选中一个下标 i 并且反转从下标 i 到下标 n - 1（包括下标 i 和下标 n - 1 ）的所有字符，成本为 n - i 。
返回使字符串内所有字符 相等 需要的 最小成本 。

反转 字符意味着：如果原来的值是 '0' ，则反转后值变为 '1' ，反之亦然。
*/
func minimumCost(s string) int64 {
	cost := int64(0)
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			// 选择较小的成本：反转前缀或反转后缀
			cost += int64(min(i, n-i))
		}
	}
	return cost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumCost("0011"))   // 2
	fmt.Println(minimumCost("010101")) // 9
}
