package main

import (
	"fmt"
	"sort"
)

/*
368. 最大整除子集
*/
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	dp := make([]int, len(nums))
	parent := make([]int, len(nums)) // 记录路径
	maxIndex := 0                    // 记录最大值的位置
	for i := range nums {
		dp[i] = 1
		parent[i] = i // 初始时每个数字指向自己
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					parent[i] = j // 记录路径
					if dp[i] > dp[maxIndex] {
						maxIndex = i
					}
				}
			}
		}
	}

	// 通过parent数组重建结果
	result := []int{}
	curr := maxIndex
	for parent[curr] != curr {
		result = append([]int{nums[curr]}, result...)
		curr = parent[curr]
	}
	result = append([]int{nums[curr]}, result...)

	return result
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3}))       // [1, 2] or [1, 3]
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8}))    // [1, 2, 4, 8]
	fmt.Println(largestDivisibleSubset([]int{4, 8, 10, 240})) // [4, 8, 240]
}
