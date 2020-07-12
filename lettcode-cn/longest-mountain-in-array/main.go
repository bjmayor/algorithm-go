package main

import (
	"fmt"
	"math"
)

//和题目trapping-rain-water 有点像
func longestMountain(A []int) int {
	var res int
	var size int
	var d int //0,无方向,1上升,-1下降。
	if len(A) < 3 {
		return 0
	}
	var pre = A[0]
	for i:=1;i<len(A);i++ {
		if A[i] > pre {
			switch d {
			case -1:
				fallthrough
			case 0:
				size = 2
			case 1:
				size++
			}
			d = 1
		} else if A[i] == pre {

			d = 0
		} else if A[i] < pre {
			switch d {
			case 0:
				size = math.MinInt32
			case 1:
				size++
			case -1:
				size++
			}
			d = -1
			if size>res {
				res = size
			}
		}
		pre = A[i]
	}
	return res
}

func main()  {
	fmt.Println(longestMountain([]int{2,1,4,7,3,2,5})) //5
	fmt.Println(longestMountain([]int{2,2,2})) //0
}
