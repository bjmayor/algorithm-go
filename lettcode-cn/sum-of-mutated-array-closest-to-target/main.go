package main

import (
	"fmt"
	"math"
	"sort"
)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	m := make(map[int]int) //保存<=arr[i]的值
	sum := make(map[int]int) //保存<=arr[i]的和
	pre := 0
	for i:=0;i<len(arr);i++ {
		m[arr[i]] = i+1
		sum[arr[i]] = pre + arr[i]
		pre = sum[arr[i]]
	}
	match := target/len(arr)
	result :=  match
	small :=  math.MaxInt32
	closeIndex := -1
	nextBig := 0
	closesmall:=match
	for match <= 100000 {
		if nextBig == 0 {
			for i:=len(arr)-1;i>=0;i-- {
				if arr[i]<=match {
					closeIndex = i
					break
				} else {
					nextBig = arr[i]
				}
			}
			if closeIndex == -1 {
				nextBig = arr[0]
			}
		}
		if match <  nextBig {
			if closeIndex != -1 {
				closesmall = arr[closeIndex]
			}
		} else {
			closesmall = nextBig
			if closeIndex+2 < len(arr) {
				nextBig = arr[closeIndex+2]
				closeIndex++
			} else {
				closeIndex= len(arr) -1
				nextBig = math.MaxInt32
			}
		}
		largerNum := len(arr)-m[closesmall]
		nextSmall := abs(sum[closesmall] + largerNum * match, target)
		if nextSmall < small {
			result = match
			small = nextSmall
			if small == 0 {
				break
			}
		} else {
			break
		}
		step := 1
		if largerNum > 0 {
			step = min(nextBig-match, nextSmall) / largerNum
		}
		if step < 1 {
			step = 1
		}
		match+= step
	}
	return result
}

func min( x, y int)  int {
	if x < y {
		return x
	}
	return y
}

func abs(x,y int)int   {
	if x > y {
		return x-y
	}
	return y-x
}

func main()  {
	tests := []struct{
		Nums []int
		Target int
		Expected int
	} {
		{ []int{15,1,1,1,1,1,1,1,1,1,1,1} ,50 ,15},
		{ []int{4,9,3} ,10 ,3},
		{ []int{2,3,5} ,10 ,5},
		{ []int{2,3,5} ,11 ,5},
		{ []int{1,2,23,24,34,36} ,110 ,30},
		{ []int{60864,25176,27249,21296,20204} ,56803 ,11361},
		{ []int{1547,83230,57084,93444,70879} ,71237,17422},
	}
	for _, t := range tests {
		if findBestValue(t.Nums, t.Target) !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", findBestValue(t.Nums, t.Target), t.Nums)
		}
	}
}