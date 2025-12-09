package main

const MOD = 1e9 + 7

func specialTriplets(nums []int) int {
	tripletCount := 0
	n := len(nums)
	freqPrev := make(map[int]int)
	freqNext := make(map[int]int)
	freqPrev[nums[0]]++
	// Initialize frequency map for the next elements
	for i := 1; i < n; i++ {
		freqNext[nums[i]]++
	}
	for j := 1; j < n-1; j++ {
		freqNext[nums[j]]--
		if freqNext[nums[j]] == 0 {
			delete(freqNext, nums[j])
		}
		target := nums[j] * 2
		if freqPrev[target] > 0 && freqNext[target] > 0 {
			contribution := (freqPrev[target] * freqNext[target]) % MOD
			tripletCount = (tripletCount + contribution) % MOD
		}
		freqPrev[nums[j]]++
	}
	return tripletCount
}
