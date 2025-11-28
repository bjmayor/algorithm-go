package main

import "fmt"

func countPalindromicSubsequence(s string) int {
	count := 0
	for c := 'a'; c <= 'z'; c++ {
		left := -1
		right := -1
		for i, ch := range s {
			if ch == c {
				if left == -1 {
					left = i
				}
				right = i
			}
		}
		if left != -1 && right > left+1 {
			seen := make(map[byte]bool)
			for i := left + 1; i < right; i++ {
				seen[s[i]] = true
			}
			count += len(seen)
		}
	}
	return count
}

func main() {
	// Test examples
	testCases := []struct {
		s      string
		expect int
	}{
		{"aabca", 3},
		{"adc", 0},
		{"bbcbaba", 4},
	}

	allPassed := true
	for _, tc := range testCases {
		result := countPalindromicSubsequence(tc.s)
		if result != tc.expect {
			fmt.Printf("Input: %s, Expected: %d, Got: %d\n", tc.s, tc.expect, result)
			allPassed = false
		}
	}

	if allPassed {
		fmt.Println("all passed")
	}
}
