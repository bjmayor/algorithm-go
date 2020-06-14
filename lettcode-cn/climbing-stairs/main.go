package main

import (
	"fmt"
	"math"
)

//递归加缓存
var cache = make(map[int]int,0 )
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	if v, ok := cache[n];ok {
		return v
	} else {
		v = climbStairs(n-2) + climbStairs(n-1)
		cache[n] = v
		return v
	}
}

//其实只存前2个就可以了，类似斐波拉契数列
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n== 2 {
		return 2
	}
	ppre := 1
	pre := 2
	for i:=3;i<=n;i++ {
		ppre, pre = pre, ppre+pre
	}
	return pre
}


//利用F(n+1) = F(n) + F(n-1) 的矩阵形式可以得到O(lgn)的时间复杂度算法
func climbStairs3(n int) int {
	return 0
}

//这是利用数学公式直接求解....有误差...
func climbStairs4(n int) int {
	sqrt5 := math.Sqrt(5)
	return int((math.Pow((1+sqrt5)/2, float64(n+1))-math.Pow((1-sqrt5)/2, float64(n+1)))/sqrt5)
}

func main()  {
	tests := []struct{
		Nums int
		Expected int
	} {
		{ 2,2},
		{ 3,3},
		{ 44,1134903170},
	}
	for _, t := range tests {
		var real = climbStairs4(t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
