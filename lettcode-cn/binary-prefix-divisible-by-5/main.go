package main

func prefixesDivBy5(nums []int) []bool {
	result := make([]bool, len(nums))
	remainder := 0
	for i, num := range nums {
		remainder = (remainder*2 + num) % 5
		result[i] = (remainder == 0)
	}
	return result
}
