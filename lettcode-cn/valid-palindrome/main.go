package main

import "fmt"

func isPalindrome(s string) bool {
	i:=0;
	j:=len(s)-1
	for i<=j {
		if v, ok := tolower(s[i]);ok {
			if vv, vok := tolower(s[j]); vok {
				if v!=vv {
					return false
				}
				i++
				j--
			} else {
				j--
			}
		} else {
			i++
		}
	}
	return true
}

func tolower(c byte) (uint8, bool) {
	if c>='A' && c<='Z' {
		return c+'a'-'A',true
	}
	if c>='0' && c<='9' {
		return c, true
	}
	if c>='a' && c<='z' {
		return c, true
	}
	return '?', false
}

func main()  {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}
