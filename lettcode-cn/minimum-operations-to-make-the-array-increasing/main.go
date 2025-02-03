package main

func minOperations(nums []int) (ans int) {
	expected := nums[0] + 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < expected {
			ans += expected - nums[i]
			expected++
		} else {
			expected = nums[i] + 1
		}
	}
	return
}
