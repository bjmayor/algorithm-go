package main

import (
	"fmt"
	"math"
)

//暴力法
func maxArea(height []int) int {
	area := math.MinInt32
	for i:=0;i<len(height)-1;i++ {
		for j:=i+1;j<len(height);j++ {
			a := (j-i) * min(height[i],height[j])
			if a> area {
				area =  a
			}
		}
	}
	return area
}

//面积 = 大的*距离-差值*距离
func maxArea2(height []int) int {
	area := 0
	left :=0
	right := len(height)-1
	for left < right {
		shorter := height[left]
		n := right-left
		if height[right] < shorter {
			shorter = height[right]
			right--
		} else {
			left++
		}
		a := n * shorter
		if a> area {
			area =  a
		}
	}
	return area
}

func min(x,y int) int  {
	if x < y {
		return x
	}
	return y
}

func main()  {
	fmt.Println(maxArea2([]int{1,8,6,2,5,4,8,3,7}))
	fmt.Println(maxArea2([]int{1,1}))
}