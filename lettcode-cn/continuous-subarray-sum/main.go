package main

import "fmt"

func main()  {
	for i, t := range tests {
		var real = checkSubarraySum(t.Nums, t.K)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real, "\t case",t.Nums, t.K)
		}
	}
}
// 暴力求解 nxn/2
func checkSubarraySum1(nums []int, k int) bool {
	for i:=0;i<len(nums)-1;i++ {
		var total int
		for j:=i;j<len(nums);j++ {
			total += nums[j]
			if j-i>=1 && total % k == 0  {
				return true
			}
		}

	}
	return false
}
/**
存下前n个数字的连续和
 */
func checkSubarraySum(nums []int, k int) bool {
	presum := make(map[int]struct{})
	presum[0] = struct{}{}
	last := 0
	pre := nums[0]%k
		for i:=1;i<=len(nums)-1;i++ {
			last, pre = pre, (pre+nums[i])%k
			if _,ok := presum[pre];ok {
				return true
			}
			presum[last] = struct{}{}
		}
	return false
}

