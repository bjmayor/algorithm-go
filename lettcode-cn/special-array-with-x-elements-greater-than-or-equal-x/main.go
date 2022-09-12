package main

import (
	"fmt"
	"sort"
)

func specialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] >= i && (i == len(nums) || nums[i] < i) {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(specialArray([]int{3, 5}))
	fmt.Println(specialArray([]int{0, 4, 3, 0, 4}))
	fmt.Println(specialArray([]int{3, 6, 7, 7, 0}))
	fmt.Println(specialArray([]int{2, 3, 0, 2}))

}
