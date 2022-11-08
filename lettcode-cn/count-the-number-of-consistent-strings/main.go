package main

func countConsistentStrings(allowed string, words []string) (res int) {
	mask := 0
	for _, c := range allowed {
		mask |= 1 << (c - 'a')
	}
	for _, word := range words {
		mask1 := 0
		for _, c := range word {
			mask1 |= 1 << (c - 'a')
		}
		if mask1|mask == mask {
			res++
		}
	}
	return
}

func main() {
	println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
}
