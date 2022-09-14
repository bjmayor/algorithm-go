package main

import "fmt"

// 就是扫描非0的，放到正确位置上。
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums) // left 指向左边处理完的尾部(都非0， right 指向处理完的头部(都是0
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	v := []int{0, 1, 0, 3, 12}
	moveZeroes(v)
	fmt.Println(v)
}
