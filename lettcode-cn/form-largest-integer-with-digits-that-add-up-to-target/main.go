package main

import (
	"fmt"
	"math"
)



//i+1的成本是cost[i], 字节字符串存储
//func largestNumber1(cost []int, target int) string {
//	dp := make([]string, target+1) // dp[i]为cost和恰好=i的值
//	dp[0] = ""
//	for i:=1;i<=target ;i++{
//		dp[i] = "-1"
//	}
//
//	for i:=8;i>=0;i--{
//		c:=cost[i]
//		for j:=c;j<=target;j++ {
//			if dp[j-c] != "-1" {
//				dp[j] = max(dp[j], dp[j-c]+strconv.Itoa(i+1)) //不用i和用i
//			}
//		}
//	}
//	if dp[target]  ==  "-1" {
//		return "0"
//	}
//	return  dp[target]
//}

//func max(a,b string) string{
//	if a[0] == '-' {
//		return b
//	}
//	if b[0] == '-' {
//		return a
//	}
//	if len(a) > len(b) {
//		return a
//	} else if len(a) < len(b) {
//		return b
//	} else {
//		for i:=0;i<len(a);i++ {
//			if a[i]>b[i]   {
//				return a
//			} else if a[i]<b[i] {
//				return b
//			}
//		}
//		return a
//	}
//}

//i+1的成本是cost[i], 优化版本。不需要字符串匹配
func largestNumber(cost []int, target int) string {
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = math.MinInt32
	}
	dp[0] = 0
	for _, c := range cost {
		for j := c; j <= target; j++ {
			dp[j] = max(dp[j], dp[j-c]+1)
		}
	}
	if dp[target] < 0 {
		return "0"
	}
	ans := make([]byte, 0, dp[target])
	for i, j := 8, target; i >= 0; i-- {
		for c := cost[i]; j >= c && dp[j] == dp[j-c]+1; j -= c {
			ans = append(ans, byte('1'+i))
		}
	}
	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main()  {
	tests := []struct{
		cost []int
		target int
		Expected string
	} {
		//{ []int{4,3,2,5,6,7,2,5,5},9, "7772"},
		//{ []int{7,6,5,5,5,6,8,7,8},12, "85"},
		{ []int{2,4,6,2,4,6,4,4,4},5, "0"},
		//{ []int{6,10,15,40,40,40,40,40,40},47, "32211"},
		//{ []int{5,6,7,3,4,6,7,4,8},29, "884444444"},
	}
	for i, t := range tests {
		var real = largestNumber(t.cost, t.target)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real )
		}
	}
}