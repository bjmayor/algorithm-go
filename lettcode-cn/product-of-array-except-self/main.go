package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums)+1)
	res[0] = 1
	for i, v := range nums {
		res[i+1] = res[i] * v
	}
	total := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= total
		total *= nums[i]
	}
	return res[0:len(nums)]
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
