package main

func isValid(s string) bool {
	pairs1 := map[uint8]uint8{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	pairs := map[uint8]uint8{
		')': '(',
		']': '[',
		'}': '{',
	}
	var a []uint8
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := pairs1[c]; ok {
			a = append(a, c)
		} else {
			if len(a) == 0 || a[len(a)-1] != pairs[c] {
				return false
			}
			a = a[:len(a)-1]
		}
	}
	return len(a) == 0
}
