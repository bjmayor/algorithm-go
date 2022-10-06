package main

import "fmt"

func getHint(secret string, guess string) string {
	a, b := 0, 0
	numa := [10]int{}
	numb := [10]int{}
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			a++
		} else {
			numa[secret[i]-'0']++
			numb[guess[i]-'0']++
		}
	}
	for i, v := range numa {
		b += min(v, numb[i])
	}
	return fmt.Sprintf("%dA%dB", a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
