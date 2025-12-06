package main

import "fmt"

const MOD = int(1e9 + 7)

// countPartitions 使用滑动窗口 + 单调队列维护区间极差，前缀和快速求区间 dp 之和。
// 时间复杂度 O(n)，空间 O(n)。
func countPartitions(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n+1)     // dp[i] 表示前 i 个元素的分割方式数
	prefix := make([]int, n+1) // prefix[i] = sum(dp[0..i])
	dp[0] = 1
	prefix[0] = 1

	// 单调队列存储下标，维护当前窗口的最小值与最大值
	minQ, maxQ := make([]int, 0, n), make([]int, 0, n)
	l := 0 // 滑动窗口左端点

	for r, v := range nums { // r 为右端点
		// 维护最大值队列（递减）
		for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] <= v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, r)

		// 维护最小值队列（递增）
		for len(minQ) > 0 && nums[minQ[len(minQ)-1]] >= v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, r)

		// 收缩左端点直到极差满足条件
		for len(maxQ) > 0 && len(minQ) > 0 && nums[maxQ[0]]-nums[minQ[0]] > k {
			if maxQ[0] == l {
				maxQ = maxQ[1:]
			}
			if minQ[0] == l {
				minQ = minQ[1:]
			}
			l++
		}

		// 有效的起点区间为 [l, r]，对应 dp 累加区间为 dp[l..r]
		// dp[r+1] = prefix[r] - prefix[l-1]
		leftPrefix := 0
		if l > 0 {
			leftPrefix = prefix[l-1]
		}
		dp[r+1] = (prefix[r] - leftPrefix + MOD) % MOD
		prefix[r+1] = (prefix[r] + dp[r+1]) % MOD
	}

	return dp[n]
}

func main() {
	// 测试用例 1
	nums1 := []int{9, 4, 1, 3, 7}
	k1 := 4
	fmt.Println(countPartitions(nums1, k1)) // 输出：6

	// 测试用例 2
	nums2 := []int{3, 3, 4}
	k2 := 0
	fmt.Println(countPartitions(nums2, k2)) // 输出：2
}
