package main

import (
	"fmt"
)

func mySqrt(x int) int {
	return mySqrt2(x)
}

//二分法
func mySqrt1(x int) int {
	left := 0
	right := x
	if x == 1 {
		return 1
	}
	for right-left>1 {
		mid := (right+left)/2
		y := mid *mid
		if y == x {
			return mid
		} else if y > x{
			right = mid
		} else {
			left = mid
		}
	}
	return left
}


//牛顿逼近法 方法使用函数f(x)的泰勒级数的前面几项来寻找方程f(y)的根。https://baike.baidu.com/item/%E7%89%9B%E9%A1%BF%E9%80%BC%E8%BF%91%E6%B3%95/1516472

func mySqrt2(x int) int {
	var result = float64(x)
	for true {
		pre :=result
		result = result/2.0+float64(x)/2.0/result
		if abs(result,pre) <= 0.01 {
			return int(result)
		}

	}
	return int(result)
}

func abs(x,y float64) float64 {
	if x >y {
		return x-y
	}
	return y-x
}

func main()  {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(1))
}