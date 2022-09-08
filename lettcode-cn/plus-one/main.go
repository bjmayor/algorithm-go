package main

func plusOne(digits []int) []int {
	plus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+plus > 9 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
				break
			}
		} else {
			digits[i] += plus
			break
		}
	}
	return digits
}
