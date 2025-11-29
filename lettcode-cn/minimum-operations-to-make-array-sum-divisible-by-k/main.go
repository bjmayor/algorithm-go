package main

import "fmt"

// 3512. 使数组和能被 K 整除的最少操作次数
//
// 给你一个整数数组 nums 和一个整数 k。你可以执行以下操作任意次：
//   - 选择一个下标 i，并将 nums[i] 替换为 nums[i] - 1。
//
// 返回使数组元素之和能被 k 整除所需的最少操作次数。
//
// 思路：每次操作会使总和减 1，因此只需要将数组和 s 的余数 r = s % k 减为 0。
// 最少的操作数就是 r（若 r == 0 则为 0）。
func minOperations(nums []int, k int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total % k
}

func main() {
	examples := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 9, 7}, 5},
		{[]int{4, 1, 3}, 4},
		{[]int{3, 2}, 6},
	}

	for _, ex := range examples {
		fmt.Printf("nums=%v k=%d => minOperations=%d\n", ex.nums, ex.k, minOperations(ex.nums, ex.k))
	}
}
