package main

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	cnt := 0
	for _, v := range s {
		m[v]++
		if m[v] == 2 {
			cnt += 2
			delete(m, v)
		}
	}
	if len(m) > 0 {
		cnt++
	}
	return cnt
}

func main() {
	println(longestPalindrome("abccccdd"))
}
