package main

// dp求解
func profitableSchemes1(n int, minProfit int, group []int, profit []int) int {
	var dp = make([][][]int, len(group)+1,len(group)+1)
	for i:=0;i<=len(group);i++ {
		dp[i] = make([][]int, n+1, n+1)
		for j:=0;j<=n;j++ {
			dp[i][j] = make([]int,minProfit+1, minProfit+1)
		}
	}
	dp[0][0][0] = 1
	var mod = 1000000007
	for i:=0;i<len(group);i++ {
		people := group[i]
		profitNow := profit[i]
		for j:=0;j<=n;j++{
			for k:=0;k<=minProfit;k++ {
				dp[i+1][j][k] = dp[i][j][k] % mod// 不选择
				idx := max(0, k-profitNow)
				if  people <= j{ //选择
					dp[i+1][j][k]	+=  dp[i][j-people][idx]
				}
			}
		}
	}
	var sum int
	for i:=0;i<=n;i++ {
		sum = (sum+dp[len(group)][i][minProfit]) % mod
	}
	return sum
}

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	g:=len(group)
	dp:=make([][]int,n+1)
	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int,minProfit+1)
	}
	for i:=0;i<n+1;i++ {
		dp[i][0]=1
	}
	for i:=0;i<g;i++ {
		manual:=group[i]
		money:=profit[i]
		for j:=n;j>=manual;j--{
			for k:=minProfit;k>=0;k-- {
				dp[j][k]+=dp[j-manual][max(0,k-money)]
				dp[j][k]%=1e9+7
			}
		}
	}
	return dp[n][minProfit]
}

func max(i,j int) int {
	if i > j {
		return i
	}
	return j
}
func main()  {
	println(profitableSchemes(5, 3, []int{2,2}, []int{2,3}))
	println(profitableSchemes(10, 5, []int{2,3,5}, []int{6,7,8}))
}
