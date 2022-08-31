package main

func shuffle(nums []int, n int) []int {
	var arr = make([]int, n*2)
	for i := 0; i < n; i++ {
		j := 2 * i
		arr[j] = nums[i]
		arr[j+1] = nums[n+i]
	}
	return arr
}
