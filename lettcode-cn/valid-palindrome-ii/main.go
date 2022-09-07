package main

func validPalindrome(s string) bool {
	return validPalindrome2(s, false)
}
func validPalindrome2(s string, deleted bool) bool {
	if len(s) <= 1 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			if deleted {
				return j-i <= 0
			} else {
				return validPalindrome2(s[i:j], true) || validPalindrome2(s[i+1:j+1], true)
			}
		}
	}
	return true
}

func main() {
	println(validPalindrome("aba"))
	println(validPalindrome("abca"))
	println(validPalindrome("abc"))
}
