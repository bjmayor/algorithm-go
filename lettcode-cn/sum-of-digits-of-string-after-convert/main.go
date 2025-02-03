package main

import "strconv"

func getLucky(s string, k int) int {
	tmp := ""
	for _, ch := range s {
		tmp += strconv.Itoa(int(ch - 'a' + 1))
	}
	s = tmp
	v := 0
	for i := 1; i <= k; i++ {
		v = 0
		for _, ch := range s {
			v += int(ch - '0')
		}
		s = strconv.Itoa(v)
	}
	return v
}
