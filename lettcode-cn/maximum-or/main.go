package main

import (
	"fmt"
)

/*
2680. 最大或值

给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 k 。每一次操作中，你可以选择一个数并将它乘 2 。

你最多可以进行 k 次操作，请你返回 nums[0] | nums[1] | ... | nums[n - 1] 的最大值。

a | b 表示两个整数 a 和 b 的 按位或 运算。
*/
func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	// 预处理后缀或
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] | nums[i]
	}

	// 前缀或
	prefix := 0
	result := 0

	// 对每个数尝试使用k次操作
	for i := 0; i < n; i++ {
		// 计算当前数使用k次操作后的值
		curr := nums[i] << k
		// 计算结果 = 前缀或 | 当前值 | 后缀或
		val := prefix | curr | suffix[i+1]
		if val > result {
			result = val
		}
		prefix |= nums[i]
	}

	return int64(result)
}

func main() {
	fmt.Println(maximumOr([]int{12, 9}, 1))   // 30
	fmt.Println(maximumOr([]int{8, 1, 2}, 2)) // 35
}
