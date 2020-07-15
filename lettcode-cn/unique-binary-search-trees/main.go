package main

import "fmt"
var m = make(map[int]int)
//递归解法
func numTrees(n int) int {
	if v, ok := m[n];ok {
		return v
	}
	var res int
	for i:=1;i<=n;i++ {
		res+= _numTress(1,i-1) * _numTress(i+1, n)
	}
	m[n] =res
	return m[n]
}

func _numTress(start, end int)  int {
	if start >=  end {
		return 1
	}
	return numTrees(end-start+1)
}

//动态规划
func numTrees2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i:=2;i<=n;i++ {
		res := 0
		for j:=1;j<=i;j++ {
			res += dp[j-1] * dp[i-j]
		}
		dp[i] = res
	}
	return dp[n]
}

//数学求解 由动态规划可以得出计算公式
// dp[0] = 1
// dp[1] = 1
// dp[2] = dp[1]+dp[1] = 2
// dp[3] = dp[2]+dp[1]*dp[1]+dp[2] = 5
// dp[4] = dp[3]+dp[2]*dp[1]+dp[1]*dp[2]+dp[3] = 14
//dp[n] = dp[j]*dp[n-j-1] //j=0,....n-1
//h(n)=(4n-2)/(n+1)*h(n-1)(n>1) h(0)=1
func numTrees3(n int) int {
	C := 1 //h(0)=1
	for i := 1; i <= n; i++ {
		C = (4*i-2)*C/(i+1)
	}
	return C
}

func main()  {
	//fmt.Println(numTrees(3)) //5
	fmt.Println(numTrees3(18)) //477638700
}


