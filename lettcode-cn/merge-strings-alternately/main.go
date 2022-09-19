package main

func mergeAlternately(word1 string, word2 string) string {
	var ans []byte = make([]byte, len(word1)+len(word2))
	var i, j, k int
	for i < len(word1) && j < len(word2) {
		ans[k] = word1[i]
		ans[k+1] = word2[j]
		i++
		j++
		k += 2
	}
	if i < len(word1) {
		copy(ans[k:], word1[i:])
	}
	if j < len(word2) {
		copy(ans[k:], word2[j:])
	}
	return string(ans)
}
