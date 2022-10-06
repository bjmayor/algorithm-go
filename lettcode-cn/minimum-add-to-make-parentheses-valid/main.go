package main

func minAddToMakeValid(s string) int {
	cnt := 0
	var ans int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
		} else if s[i] == ')' {
			if cnt > 0 {
				cnt--
			} else {
				ans++
			}
		}
	}
	return ans + cnt
}
