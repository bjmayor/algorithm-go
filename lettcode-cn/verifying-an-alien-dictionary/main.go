package main

func isAlienSorted(words []string, order string) bool {
	m := map[byte]int{}
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}
	for i := 1; i < len(words); i++ {
		word1 := words[i-1]
		word2 := words[i]
		j := 0
		for ; j < len(word1) && j < len(word2); j++ {
			if m[word1[j]] > m[word2[j]] {
				return false
			} else if m[word1[j]] < m[word2[j]] {
				break
			}
		}
		if j == len(word1) || j == len(word2) {
			if len(word1) > len(word2) {
				return false
			}
		}

	}
	return true
}

func main() {
	println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"))
	println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))
}
