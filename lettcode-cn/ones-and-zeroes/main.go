package main

import (
	"fmt"
)

//未压缩的原始方法
func findMaxForm1(strs []string, m int, n int) int {
	counter:=make([][2]int, len(strs), len(strs))
	for i:=0;i<len(strs);i++ {
		for j:=0;j<len(strs[i]);j++ {
			counter[i][strs[i][j]-'0']++
		}
	}
	var dp = make([][][]int, len(strs)+1, len(strs)+1)
	for i:=0;i<=len(strs);i++ {
		dp[i] = makexy(m+1,n+1)
	}
	for i:=0;i<=m;i++ {
		for j:=0;j<=n;j++ {
			for k:=0;k<len(strs);k++ {
				dp[k+1][i][j] = dp[k][i][j]
				if counter[k][0]<=i  && counter[k][1]<=j {
					dp[k+1][i][j] = max(dp[k][i][j], dp[k][i-counter[k][0]][j-counter[k][1]]+1)
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}

func findMaxForm(strs []string, m int, n int) int {
	var dp = make([][]int, m+1,m+1)
	for i:=0;i<=m;i++ {
		dp[i] = make([]int, n+1,n+1)
	}
	for k:=0;k<len(strs);k++ {
		zeros:=0
		ones:=0
		for i:=0;i<len(strs[k]);i++ {
			if strs[k][i] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		for i:=m;i>=zeros;i-- {
			for j:=n;j>=ones;j-- {
				if zeros<=i  && ones<=j {
					dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
				}
			}
		}

	}


	return dp[m][n]
}




func max(x,y int) int  {
	if x >y {
		return x
	}
	return y
}
func makexy(x,y int) [][]int  {
	var result = make([][]int, x,x)
	for i:=0;i<x;i++ {
		result[i] = make([]int, y,y)
	}
	return result
}

func main()  {
	for i, t:= range tests	{
		real := findMaxForm(t.Strs, t.M, t.N)
		if real != t.Expected {
			fmt.Printf("case %d: expected:%d, real:%d\n", i, t.Expected, real)
		}
	}
}