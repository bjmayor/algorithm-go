package main

import "strings"

func stringMatching(words []string) []string {
	var res []string
	for _, word1 := range words {
		for _, word2 := range words {
			if word1 != word2 {
				if strings.Contains(word2, word1) {
					res = append(res, word1)
					break
				}
			}
		}
	}
	return res
}
