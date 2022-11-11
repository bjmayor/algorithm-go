package main

func halvesAreAlike(s string) bool {
	cnt := 0
	for i := 0; i < len(s)/2; i++ {
		if isyuanyin(s[i]) {
			cnt++
		}
	}
	for i := len(s) / 2; i < len(s); i++ {
		if isyuanyin(s[i]) {
			cnt--
		}
	}
	return cnt == 0
}

func isyuanyin(v uint8) bool {
	arr := []uint8{
		'a', 'o', 'e', 'i', 'u',
		'A', 'O', 'E', 'I', 'U',
	}
	for _, c := range arr {
		if v == c {
			return true
		}
	}
	return false
}
