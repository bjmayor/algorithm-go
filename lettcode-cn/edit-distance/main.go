package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1, len(word1)+1) //word1前i个字符 和 word2前j个字符变更为同样的单词的最少步骤数
	for i:=0;i<len(word1)+1;i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	dp[0][0] = 0
	for i:=1;i<=len(word1);i++ {
		dp[i][0] = i//都是删除
	}
	for i:=1;i<=len(word2);i++ {
		dp[0][i] = i//都是插入
	}
	for i:=1;i<=len(word1);i++ {
		for j:=1;j<=len(word2);j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])+1//变过来的方式有 插入，替换，删除
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func min(values ...int) int  {
	v := values[0]
	for _, vv := range values {
		if vv < v {
			v = vv
		}
	}
	return v
}

func main()  {
	tests := []struct {
		word1 string
		word2 string
		Expected int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}
	for _, t := range tests {
		var real = minDistance(t.word1, t.word2)
		if real != t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, "data:", t.word1, t.word2)
		}
	}
}