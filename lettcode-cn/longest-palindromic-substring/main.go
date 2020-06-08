package main

import (
	"fmt"
	"strings"
)

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
	if len(s) < 2 || (len(s) == 2 && s[0] == s[1]) {
		return s
	}
	var p = make([]int, len(s)*2+1)
	p[0] = 1
	p[1] = 3
	var i = 2
	var maxRight = 2//最右边
	var center = 1
	s = strings.Join(strings.Split(s,""),"#")
	s = "#"+s+"#"
	max := 3
	var result = s[0:3]
	//选定i是中心点,往2边扩
	for i=2;i<len(s); i++{
		var mirror = 2 * center -i
		if i < maxRight {
			p[i] = min(p[mirror], maxRight-i)
		} else {
			p[i] = 1
		}
		center = i
		maxRight = i+p[i]/2
		for maxRight+1<len(s) &&  (2*center-maxRight-1>=0) && s[maxRight+1] == s[2*center-maxRight-1] {
			maxRight++
		}
		p[i] = (maxRight-center)*2+1
		if p[i] > max {
		  max = p[i]
		  result = s[2*center-maxRight:maxRight+1]
		}
		if maxRight == len(s)-1 {
			break
		}
	}
	var r string
	for i, _ := range result {
		if i % 2 == 1 {
			r += result[i:i+1]
		}
	}
	return r
}

func min(x, y int) int  {
	if x < y {
		return x
	}
	return y
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
