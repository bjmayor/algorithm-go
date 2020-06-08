package main

import "fmt"

func isPalindrome(s string) bool  {
	for i:=0;i<=len(s)/2;i++{
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
func longestPalindrome(s string) string {
	return longestPalindrome2(s)
}

//暴力破解法
func longestPalindrome1(s string) string {
	max := 0
	var result string
	for i:=0;i<len(s);i++ {
		for j:=i;j<len(s);j++ {
			if isPalindrome(s[i:j+1])  && j-i+1 > max{
				max = j-i+1
				result = s[i:j+1]
			}
		}
	}
	return result
}

//动态规划法
func longestPalindrome2(s string) string {
	dp := make([][]bool, len(s), len(s))

	for i:=0;i<len(s);i++ {
		dp[i] = make([]bool, len(s), len(s))
		dp[i][i] = true
	}

	for i:=len(s)-2;i>=0;i--{
		for j:=len(s)-1;j>i;j-- {
			if s[i] == s[j] {
				if j - i  <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
		}
	}
	var max int
	var result  string
	for i:=0;i<len(s);i++ {
		for j:=i;j<len(s);j++ {
			if dp[i][j] && j-i+1 > max {
				max = j-i+1
				result = s[i:j+1]
			}
		}
	}
	return result
}

//中心扩展算法
func longestPalindrome3(s string) string {
	max := 0
	var result string
	//选定i是中心点,往2边扩
	for i:=0;i<len(s); i++{
		end :=i
		for j:=i+1;j<len(s);j++ {
			if s[j] != s[i] {
				break
			} else {
				end = j
			}
		}
		if max < end-i+1{
			max = end-i+1
			result = s[i:end+1]
		}
		var j = i
		var k= end
		for ;j<len(s) && k>=0; {
			if s[k] == s[j] {
				k--
				j++
			} else {
				break
			}
		}
		if j-k-1 > max {
			max = j-k-1
			result = s[k+1:j]
		}
	}
	return result
}

//Manacher 算法
func longestPalindrome4(s string) string {
	max := 0
	var result string
	//选定i是中心点,往2边扩
	for i:=0;i<len(s); i++{
		end :=i
		for j:=i+1;j<len(s);j++ {
			if s[j] != s[i] {
				break
			} else {
				end = j
			}
		}
		if max < end-i+1{
			max = end-i+1
			result = s[i:end+1]
		}
		var j = i
		var k= end
		for ;j<len(s) && k>=0; {
			if s[k] == s[j] {
				k--
				j++
			} else {
				break
			}
		}
		if j-k-1 > max {
			max = j-k-1
			result = s[k+1:j]
		}
		i = (j+k)/2
	}
	return result
}
func main() {
	tests := []struct {
		Nums     string
		Expected string
	}{
		{"a", "a"},
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"ccc", "ccc"},
		{"ccd", "cc"},
		{"abcda", "a"},
		{"aaabaaaa", "aaabaaa"},
	}
	for _, t := range tests {
		var real = longestPalindrome4(t.Nums)
		if real != t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real)
		} else {
			fmt.Println("ok")
		}
	}
}
