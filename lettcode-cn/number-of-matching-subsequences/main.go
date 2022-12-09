package main

func numMatchingSubseq(s string, words []string) int {
	var ans int
	var jmp [][26]int = make([][26]int, len(s))
	for i := 0; i < len(jmp); i++ {
		for j := 0; j < len(jmp[i]); j++ {
			jmp[i][j] = -1
		}
	}
	for i := len(s) - 2; i >= 0; i-- {
		jmp[i][s[i+1]-'a'] = i + 1
		for j := 0; j < len(jmp[i+1]); j++ {
			if j != int(s[i+1]-'a') {
				jmp[i][j] = jmp[i+1][j]
			}
		}
	}
	for _, word := range words {
		if isSubSeq(jmp, s, word) {
			ans++
		}
	}
	return ans
}

func isSubSeq(jmp [][26]int, full, sub string) bool {
	if len(sub) > len(full) {
		return false
	}
	i, j := 0, 0
	for i < len(full) && j < len(sub) {
		for sub[j] == full[i] {
			j++
			if j == len(sub) {
				return true
			}
			i = jmp[i][sub[j]-'a']
			if i == -1 {
				return false
			}
		}
		i++
	}
	return false
}

func main() {
	println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
}
