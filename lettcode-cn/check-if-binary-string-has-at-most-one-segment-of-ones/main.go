package main

func checkOnesSegment(s string) bool {
	left, right := 0, len(s)-1
	for ; left < len(s); left++ {
		if s[left] == '0' {
			break
		}
	}

	for ; right >= 0; right-- {
		if s[right] == '1' {
			break
		}
	}
	return left-right == 1
}
