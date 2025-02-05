package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	for even < len(nums) && odd < len(nums) {
		if nums[even]%2 == 0 {
			even += 2
		}
		if nums[odd]%2 == 1 {
			odd += 2
		}
		if even < len(nums) && odd < len(nums) && nums[even]%2 == 1 && nums[odd]%2 == 0 {
			nums[even], nums[odd] = nums[odd], nums[even]
		}
	}
	return nums
}

func main() {
	nums := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(nums))

	nums2 := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParityII(nums2))
}
