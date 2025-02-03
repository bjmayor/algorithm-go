package main

func checkIfPangram2(sentence string) bool {
	var cnt [26]int
	for _, ch := range sentence {
		cnt[ch-'a'] = 1
	}
	for _, v := range cnt {
		if v == 0 {
			return false
		}
	}
	return true
}

func checkIfPangram(sentence string) bool {
	state := 0
	for _, c := range sentence {
		state |= 1 << (c - 'a')
	}
	return state == 1<<26-1
}
