package main

func runningSum(nums []int) []int {
	var ans = make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			ans[0] = nums[0]
		} else {

			ans[i] += ans[i-1] + v
		}
	}
	return ans
}
