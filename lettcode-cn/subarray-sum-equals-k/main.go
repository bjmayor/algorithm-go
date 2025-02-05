package main

import "fmt"

func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := 0
	prefixSumCount := make(map[int]int)
	prefixSumCount[0] = 1

	for _, num := range nums {
		prefixSum += num
		if c, found := prefixSumCount[prefixSum-k]; found {
			count += c
		}
		prefixSumCount[prefixSum]++
	}

	return count
}

func main() {
	// 示例 1
	nums1 := []int{1, 1, 1}
	k1 := 2
	fmt.Println(subarraySum(nums1, k1)) // 输出: 2

	// 示例 2
	nums2 := []int{1, 2, 3}
	k2 := 3
	fmt.Println(subarraySum(nums2, k2)) // 输出: 2
}
