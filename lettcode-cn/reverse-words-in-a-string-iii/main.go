package main

func reverseWords(s string) string {
	start := 0
	t := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			for m, n := start, i; m < (i+start)/2; m++ {
				t[m], t[start+n-m-1] = t[start+n-m-1], t[m]
			}
			start = i + 1
		}
	}
	for m, n := start, len(s); m < (len(s)+start)/2; m++ {
		t[m], t[start+n-m-1] = t[start+n-m-1], t[m]
	}
	return string(t)
}

func main() {
	println(reverseWords("Let's take LeetCode contest"))
}
