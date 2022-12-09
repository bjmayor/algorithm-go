package main

func minOperations(s string) int {
	n := len(s)
	cnt := 0
	for i := 0; i < n; i++ {
		if int(s[i]-'0') != i%2 {
			cnt++
		}
	}
	if cnt < n-cnt {
		return cnt
	}
	return n - cnt
}
