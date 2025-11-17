package main

func kLengthApart(nums []int, k int) bool {
	hasOne := false
	cnt := 0
	for _, v := range nums {
		if v == 1 {
			if hasOne {
				if cnt < k {
					return false
				}
			}
			hasOne = true
			cnt = 0
		} else {
			cnt++
		}
	}
	return true
}
