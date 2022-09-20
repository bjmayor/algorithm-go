package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	all := 0
	for _, v := range nums {
		all += v
	}
	if all%k > 0 {
		return false
	}
	per := all / k
	sort.Ints(nums)
	n := len(nums)
	if nums[n-1] > per {
		return false
	}

	dp := make([]bool, 1<<n)
	for i := range dp {
		dp[i] = true
	}
	var dfs func(int, int) bool
	dfs = func(s, p int) bool {
		if s == 0 {
			return true
		}
		if !dp[s] {
			return dp[s]
		}
		dp[s] = false
		for i, num := range nums {
			if num+p > per {
				break
			}
			if s>>i&1 > 0 && dfs(s^1<<i, (p+nums[i])%per) {
				return true
			}
		}
		return false
	}
	return dfs(1<<n-1, 0)
}

func main() {
	//println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	//println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
	println(canPartitionKSubsets([]int{4, 4, 6, 2, 3, 8, 10, 2, 10, 7}, 4))
}
