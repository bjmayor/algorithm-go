package main

import (
	"fmt"
	"math"
)

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)

	// 计算前缀和
	sum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(nums[i])
	}

	// 按模分组，记录每组的最小前缀和
	minSum := make([]int64, k)
	for i := range minSum {
		minSum[i] = math.MaxInt64
	}
	minSum[0] = 0 // 空数组的前缀和为 0

	ans := int64(math.MinInt64)

	for j := 0; j < n; j++ {
		mod := (j + 1) % k
		// 用当前位置减去同组的最小前缀和
		if minSum[mod] != math.MaxInt64 {
			ans = max(ans, sum[j+1]-minSum[mod])
		}
		// 更新该组的最小前缀和
		minSum[mod] = min(minSum[mod], sum[j+1])
	}

	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int64
	}{
		{
			name:     "测试用例1：基础情况",
			nums:     []int{1, 2},
			k:        1,
			expected: 3,
		},
		{
			name:     "测试用例2：全是负数",
			nums:     []int{-1, -2, -3, -4, -5},
			k:        4,
			expected: -10,
		},
		{
			name:     "测试用例3：正负混合",
			nums:     []int{-5, 1, 2, -3, 4},
			k:        2,
			expected: 4,
		},
		{
			name:     "测试用例4：k 等于数组长度",
			nums:     []int{1, 2, 3, 4, 5},
			k:        5,
			expected: 15,
		},
		{
			name:     "测试用例5：需要选择中间某段",
			nums:     []int{-1, 10, -2, 8, -3},
			k:        2,
			expected: 15, // 修正期望值：[-1, 10, -2, 8] = 15
		},
		{
			name:     "测试用例6：大数测试",
			nums:     []int{1000000000, -1000000000, 1000000000},
			k:        3,
			expected: 1000000000,
		},
		{
			name:     "测试用例7：k=1，所有子数组都满足",
			nums:     []int{-2, 5, -1, 7, -3},
			k:        1,
			expected: 11,
		},
	}

	allPassed := true
	for _, tc := range tests {
		result := maxSubarraySum(tc.nums, tc.k)
		if result != tc.expected {
			fmt.Printf("FAIL: %s\n", tc.name)
			fmt.Printf("输入: nums = %v, k = %d\n", tc.nums, tc.k)
			fmt.Printf("输出: %d\n", result)
			fmt.Printf("期望: %d\n\n", tc.expected)
			allPassed = false
		}
	}

	if allPassed {
		fmt.Println("All passed")
	}
}
