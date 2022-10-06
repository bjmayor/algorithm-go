package main

func characterReplacement(s string, k int) int {
	m := map[byte]int{}
	var ans = 0
	start, cur, maxCount := 0, 0, 0

	for i := 0; i < len(s); i++ {
		m[s[i]]++
		maxCount = max(maxCount, m[s[i]])
		cur++
		for maxCount+k < cur {
			m[s[start]]--
			start++
			cur--
		}
		if cur > ans {
			ans = cur
		}

	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	//println(characterReplacement("ABAB", 2))
	println(characterReplacement("AABABBA", 1))
}
