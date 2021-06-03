package main

import (
	"fmt"
	"math"
)
/**
1. 傻瓜方法就是从sqrt(num+2)开始尝试到1。
2. 均值不等式思想, 均值不等式可知 a+b>=2*sqrt(a*b)，两个数和固定的情况下，两数越接近，他们的乘积越大。
证明:
设其中一个数为x,两数之和为一常数a
则另一个数为a-x
x(a-x)=ax-x²=-(x²-ax+a²/4)+a²/4=a²/4-(x-a/2)²
上式越大则要求x-a/2越小.即x越接近a/2.
因此两个数越接近,乘积越大
 */
func closestDivisors(num int) []int {
	var m,n int;
	n = int(math.Ceil(math.Sqrt(float64(num))))
	m = n;
	for {
		if m*n == num+1 || m*n == num+2 {
			return []int{m,n}
		} else if (m*n < num+1) {
			n++//n是大数
		} else {
			m--;//m是小数
		}
	}
}

func main()  {
	tests := []struct{
		Num int
		ExpectedX int
		ExpectedY int
	} {
		{ 1,1,2},
		{ 8,3,3},
		{ 123,5,25},
		{ 999,40,25},
	}
	for i, t := range tests {
		var real = closestDivisors(t.Num)
		if real[0]*real[1] !=  t.ExpectedX*t.ExpectedY{
			fmt.Println(i, " expected:", t.ExpectedX, t.ExpectedY, " real:", real, t.Num)
		}
	}
}