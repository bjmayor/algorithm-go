package main

func longestPalindrome(words []string) int {
	var ans int
	m := map[string]int{}
	for _, word := range words {
		other := word[1:] + word[0:1]
		if _, ok := m[other]; ok {
			ans += 4
			m[other]--
			if m[other] == 0 {
				delete(m, other)
			}
		} else {
			m[word]++
		}
	}
	for k, _ := range m {
		if k[0] == k[1] {
			ans += 2
			break
		}
	}
	return ans
}

func main() {
	println(longestPalindrome([]string{"lc", "cl", "gg"}))                                                                   // 6
	println(longestPalindrome([]string{"ab", "ty", "yt", "lc", "cl", "ab"}))                                                 // 8
	println(longestPalindrome([]string{"cc", "ll", "xx"}))                                                                   //2
	println(longestPalindrome([]string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"})) // 22
}
