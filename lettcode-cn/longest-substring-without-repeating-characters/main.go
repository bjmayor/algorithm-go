package main

import "fmt"



func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var  start, end int
	end = 1
	longest := 1
	for i:=1; i< len(s);i++ {
		found := false
		var j int
		for j =start;j<end;j++ {
			if s[j] == s[i] {
				found = true
				break
			}
		}
		if found {
			if s[i] == s[start] {
				start++
			} else {
				if  end-start > longest {
					longest = end-start
				}
				start = j+1
			}
		}
		end++
	}
	if  end-start > longest {
		longest = end-start
	}
	return longest
}

func main() {
	tests := []struct{
		Nums string
		Expected int
	} {
		{ "abcabcbb" ,3 },
		{ "bbbbb" ,1 },
		{ "pwwkew" ,3 },
		{ "" ,0 },
	}
	for _, t := range tests {
		if lengthOfLongestSubstring(t.Nums) !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", lengthOfLongestSubstring(t.Nums), t.Nums)
		}
	}
}
