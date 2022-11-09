package main

import "math"

func minWindow(s string, t string) string {
	m := map[uint8]int{}
	target := map[uint8]int{}
	for i := 0; i < len(t); i++ {
		target[t[i]]++
	}
	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && target[s[r]] > 0 {
			m[s[r]]++
		}
		for contains(m, target) && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := target[s[l]]; ok {
				m[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

func contains(m, target map[uint8]int) bool {
	for k, v := range target {
		if m[k] < v {
			return false
		}
	}
	return true
}

func main() {
	println(minWindow("ADOBECODEBANC", "ABC"))
}
