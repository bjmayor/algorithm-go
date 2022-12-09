package main

func numSubarrayBoundedMax(nums []int, left int, right int) (res int) {
	last2, last1 := -1, -1 // last1 在区间内
	for i, x := range nums {
		if left <= x && x <= right {
			last1 = i
		} else if x > right {
			last2 = i
			last1 = -1
		}
		if last1 != -1 {
			res += last1 - last2 // 以last1为右端点的子数组个数
		}
	}
	return
}
