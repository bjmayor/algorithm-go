package main

import (
	"fmt"
)

/*
3305. 元音辅音字符串计数 I

给你一个字符串 word 和一个 非负 整数 k。

返回 word 的 子字符串 中，每个元音字母（'a'、'e'、'i'、'o'、'u'）至少 出现一次，并且 恰好 包含 k 个辅音字母的子字符串的总数。
*/

func countOfSubstrings(word string, k int) int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	count := func(m int) int {
		n := len(word)
		var res int = 0
		consonants := 0
		occur := make(map[byte]int)
		for i, j := 0, 0; i < n; i++ {
			for j < n && (consonants < m || len(occur) < 5) {
				if vowels[word[j]] {
					occur[word[j]]++
				} else {
					consonants++
				}
				j++
			}
			if consonants >= m && len(occur) == 5 {
				res += n - j + 1
			}
			if vowels[word[i]] {
				occur[word[i]]--
				if occur[word[i]] == 0 {
					delete(occur, word[i])
				}
			} else {
				consonants--
			}
		}
		return res
	}
	return count(k) - count(k+1)
}

func main() {
	fmt.Println(countOfSubstrings("aeioqq", 1))         // expected 0
	fmt.Println(countOfSubstrings("aeiou", 0))          // expected 1
	fmt.Println(countOfSubstrings("ieaouqqieaouqq", 1)) // expected 3
}
