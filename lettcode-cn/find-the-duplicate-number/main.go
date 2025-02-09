package main

import "fmt"

func findDuplicate(nums []int) int {
	last := nums[len(nums)-1]
	for last != nums[last-1] {
		temp := nums[last-1]
		nums[last-1] = last
		last = temp
	}
	return last
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
