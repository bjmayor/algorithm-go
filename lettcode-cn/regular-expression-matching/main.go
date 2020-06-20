package main

import "fmt"

func isMatch(s string, p string) bool {
	return isMatch2(s,p)
}
//递归版本
func isMatch1(s string, p string) bool {
	if  len(p) == 0  {
		return len(s) == 0
	}
	firstMatch := len(s) > 0  && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return  isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:],p))
	} else {
		return firstMatch && isMatch(s[1:],p[1:])
	}
}

//动态规划
//如果 p.charAt(j) == s.charAt(i) : dp[i][j] = dp[i-1][j-1]；
//如果 p.charAt(j) == '.' : dp[i][j] = dp[i-1][j-1]；
//如果 p.charAt(j) == '*'：
	//如果 p.charAt(j-1) != s.charAt(i) : dp[i][j] = dp[i][j-2] //in this case, a* only counts as empty
	//如果 p.charAt(i-1) == s.charAt(i) or p.charAt(i-1) == '.'：
//dp[i][j] = dp[i-1][j] //in this case, a* counts as multiple a
//or dp[i][j] = dp[i][j-1] // in this case, a* counts as single a
//or dp[i][j] = dp[i][j-2] // in this case, a* counts as empty
func isMatch2(s string, p string) bool {
	dp:=make([][]bool,len(s)+1) //dp[i][j] 表示 s 的前 i 个是否能被 p 的前 j 个匹配
	for i:=0;i<len(s)+1;i++{
		dp[i]=make([]bool,len(p)+1)
	}
	dp[0][0]=true
	for i:=2;i<len(p)+1;i++{
		if p[i-1]=='*'{
			dp[0][i]=dp[0][i-2] //?* 可以忽略
		}

	}

	for i:=1;i<len(s)+1;i++ {
		for j:=1;j<len(p)+1;j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			}else if p[j-1] == '*' {
				if dp[i][j-2] == true {
					dp[i][j] = true
				} else {
					if p[j-2] == '.' || p[j-2] == s[i-1] {
						dp[i][j] = dp[i-1][j]
					}
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

func main() {
	tests := []struct{
		Nums string
		Target string
		Expected bool
	} {
		{ "ab",".*c", false},
		{ "aaca","ab*a*c*a", true},
		{ "aaa","ab*a*c*a", true},
		{ "aaa","a*a", true},
		{ "aa","a", false},
		{ "aa","a*", true},
		{ "ab",".*", true},
		{ "aab","c*a*b", true},
		{ "mississippi","mis*is*p*.", false},
	}
	for i, t := range tests {
		if isMatch(t.Nums,t.Target) !=  t.Expected {
			fmt.Println(i," expected:", t.Expected, " real:", isMatch(t.Nums, t.Target))
		}
	}
}