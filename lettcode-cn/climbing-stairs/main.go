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
	var square = [2][2]int{{1,1},{1,0}}
	var res = pow(square, n)
	return res[1][0]+res[1][1]
}

func pow(square [2][2]int, n int ) [2][2]int  {
	var res = [2][2]int{{1,0},{0,1}} //A*B = B
	for n > 0 {
		if n & 1 == 1 {
			res = multiply(res, square)
		}
		square = multiply(square,square)
		n = n>>1
	}
	return res
}

func multiply(s1 [2][2]int, s2 [2][2]int) [2][2]int {
	var res [2][2]int
	for i:=0;i<2;i++ {
		for j:=0;j<2;j++ {
			res[i][j] = s1[i][0]*s2[0][j]+s1[i][1]*s2[1][j]
		}
	}
	return res
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
		var real = climbStairs3(t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
