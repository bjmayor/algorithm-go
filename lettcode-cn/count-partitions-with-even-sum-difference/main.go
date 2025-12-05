package main

func countPartitions(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 0 {
		return len(nums) - 1
	}
	return 0
}

func countPartitionsDiff(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	
	count := 0
	prefixSum := 0
	
	// 遍历分割点 i 从 0 到 n-2
	for i := 0; i < len(nums)-1; i++ {
		prefixSum += nums[i]
		rightSum := sum - prefixSum
		diff := prefixSum - rightSum
		
		// 判断差值是否为偶数
		if diff%2 == 0 {
			count++
		}
	}
	
	return count
}
