package main

import "fmt"

/*
*
416. 分割等和子集
*/
func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	if max > target {
		return false
	}

	// 初始化动态规划数组
	// dp[i][j] 表示前i个元素中是否存在和为j的子集
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	// 初始化状态：空集的和为0
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	// 初始化状态：第一个元素的和为nums[0]
	dp[0][nums[0]] = true
	// 动态规划转移
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j >= v {
				// 转移公式：dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
				// dp[i-1][j] 表示不选第i个元素
				// dp[i-1][j-v] 表示选第i个元素
				// 只要有一个为true，dp[i][j]就为true
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				// 如果j < v，则不能选第i个元素
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
	nums = []int{1, 2, 3, 5}
	fmt.Println(canPartition(nums))
}
