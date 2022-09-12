package main

func arraySign(nums []int) int {
	p := 0
	n := 0
	for _, v := range nums {
		if v == 0 {
			return 0
		} else if v > 0 {
			p++
		} else {
			n++
		}
	}
	if n%2 == 1 {
		return -1
	}
	return 1
}
