package main

func partitionDisjoint(nums []int) int {
	n := len(nums)
	leftMax, leftPos, curMax := nums[0], 0, nums[0]
	for i := 1; i < n-1; i++ {
		curMax = max(curMax, nums[i])
		if nums[i] < leftMax {
			leftMax = curMax
			leftPos = i
		}
	}
	return leftPos + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	//println(partitionDisjoint([]int{5, 0, 3, 8, 6}))
	//println(partitionDisjoint([]int{1, 1, 1, 0, 6, 12}))
	println(partitionDisjoint([]int{1, 1, 1}))
	println(partitionDisjoint([]int{26, 51, 40, 58, 42, 76, 30, 48, 79, 91})) // 1
}
