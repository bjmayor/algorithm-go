package main

func isPrefixOfWord(sentence string, searchWord string) int {
	words := 1
	for i := 0; i < len(sentence); i++ {
		ok := true
		for j := 0; j < len(searchWord); j++ {
			if sentence[i+j] != searchWord[j] {
				ok = false
				for i += j; i < len(sentence); i++ {
					if sentence[i] == ' ' {
						words += 1
						break
					}
				}
				break
			}
		}
		if ok {
			return words
		}
	}
	return -1
}

func main() {
	println(isPrefixOfWord("i love eating burger", "burg"))           // 4
	println(isPrefixOfWord("this problem is an easy problem", "pro")) //2
	println(isPrefixOfWord("i am tired", "you"))                      // -1
	println(isPrefixOfWord("hellohello hellohellohello", "ell"))      // -1
}
