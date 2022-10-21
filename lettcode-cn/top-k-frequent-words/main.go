package main

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	m := map[string]int{}
	for _, w := range words {
		m[w]++
	}
	uniqueWords := make([]string, 0, len(m))
	for w := range m {
		uniqueWords = append(uniqueWords, w)
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return m[s] > m[t] || m[s] == m[t] && s < t

	})
	return uniqueWords[:k]
}

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}
