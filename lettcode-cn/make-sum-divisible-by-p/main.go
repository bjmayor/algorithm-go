package main

import "fmt"

// minSubarray 返回需要移除的最短子数组长度，使得剩余元素之和能被 p 整除。
// 如果无法满足要求（或必须移除整个数组），返回 -1。
// 算法：前缀和取模 + 哈希表（余数 -> 最近索引），时间 O(n)，空间 O(min(n,p))。
func minSubarray(nums []int, p int) int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total = (total + v) % p
	}
	q := total % p
	if q == 0 {
		return 0
	}

	last := make(map[int]int)
	last[0] = -1 // 前缀为空时余数为 0，对应索引 -1

	cur := 0
	ans := n + 1
	for i, v := range nums {
		cur = (cur + v) % p
		need := (cur - q + p) % p
		if j, ok := last[need]; ok {
			// 能移除的子数组为 (j+1..i)
			if i-j < ans {
				ans = i - j
			}
		}
		// 记录当前余数对应的最近索引（覆盖为最近，以便得到更短的长度）
		last[cur] = i
	}

	if ans == n+1 || ans == n {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minSubarray([]int{3, 1, 4, 2}, 6)) // expect 1
}
