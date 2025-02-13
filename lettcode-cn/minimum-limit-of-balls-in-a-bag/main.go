package main

func minimumSize(nums []int, maxOperations int) int {
	left := 1
	right := 1
	for _, v := range nums {
		right = max(right, v)
	}
	for left < right {
		mid := (left + right) >> 1
		op := 0
		for _, v := range nums {
			if v > mid {
				op += (v - 1) / mid
			}
		}
		if op <= maxOperations {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	// 测试用例1: nums = [9], maxOperations = 2
	nums1 := []int{9}
	maxOperations1 := 2
	result1 := minimumSize(nums1, maxOperations1)
	println("测试用例1结果:", result1) // 预期输出: 3

	// 测试用例2: nums = [2,4,8,2], maxOperations = 4
	nums2 := []int{2, 4, 8, 2}
	maxOperations2 := 4
	result2 := minimumSize(nums2, maxOperations2)
	println("测试用例2结果:", result2) // 预期输出: 2

	// 测试用例3: nums = [7,17], maxOperations = 2
	nums3 := []int{7, 17}
	maxOperations3 := 2
	result3 := minimumSize(nums3, maxOperations3)
	println("测试用例3结果:", result3) // 预期输出: 7
}
