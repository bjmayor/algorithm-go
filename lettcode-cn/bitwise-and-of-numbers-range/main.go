package main

import "fmt"
func rangeBitwiseAnd(m int, n int) int {
	return rangeBitwiseAnd4(m,n)
}
//暴力解法
func rangeBitwiseAnd1(m int, n int) int {
	var r =m
	for i:=m+1;i<=n;i++ {
		r= r&i
		if r == 0 {
			return 0
		}
	}
	return r
}

//最低位解法， 最低位只要出现一个0则为0
func rangeBitwiseAnd2(m int, n int) int {
	zero := 0
	for m<n {
		m = m>>1
		n = n>>1
		zero ++
	}
	return m <<zero
}

//最低位解法，n的右边的全为0的和其他数相与必为0，
func rangeBitwiseAnd3(m int, n int) int {
	for m<n {
		n = n&(n-1)
	}
	return n
}

//暴力解法
func rangeBitwiseAnd4(m int, n int) int {
	i := m ^ n;
	i |= i >> 1;
	i |= i >> 2;
	i |= i >> 4;
	i |= i >> 8;
	i |= i >> 16;
	return  n & ^i
}

func main()  {
	tests := []struct{
		M int
		N int
		Expected int
	} {
		{ 5,7,4},
		{ 0,1,0},
		{ 4000000,2147483646,0},
	}
	for i, t := range tests {
		var real = rangeBitwiseAnd(t.M, t.N)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real)
		}
	}
}
