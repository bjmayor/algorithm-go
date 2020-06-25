package main

import (
	"fmt"
	"math"
	"sort"
)
func threeSumClosest(nums []int, target int) int {
	return threeSumClosest2(nums,target)
}
//暴力解法
func threeSumClosest1(nums []int, target int) int {
	var diff = math.MaxInt32
	var sum int
	for i:=0;i<len(nums);i++ {
		for j:=i+1;j<len(nums);j++ {
			for k:=j+1;k<len(nums);k++ {
				if abs(nums[i]+nums[j]+nums[k],target) < diff {
					diff = abs(nums[i]+nums[j]+nums[k],target)
					sum = nums[i]+nums[j]+nums[k]
				}
			}
		}
	}
	return sum
}

func abs(x,y int) int  {
	 if x > y {
	 	return x-y
	 }
	 return y-x
}

//双指针
func threeSumClosest2(nums []int, target int) int {
	sort.Ints(nums)
	diff := math.MaxInt32
	sum := 0
	for i:=0;i<len(nums)-2;i++ {
		left := i+1
		right := len(nums)-1
		ntarget := target-nums[i]
		for left < right {
			isum := nums[left] + nums[right]
			ndiff := abs(nums[i]+isum, target)
			if ndiff<diff {
				sum = nums[i]+isum
				diff = ndiff
			}
			if isum > ntarget {
				right--
			} else if  isum < ntarget {
				left++
			} else {
				return target
			}


		}
	}
	return sum
}

func main()  {
	fmt.Println(threeSumClosest([]int{1,2,4,8,16,32,64,128}, 82))	//82
	fmt.Println(threeSumClosest([]int{0,2,1,-3}, 1))	//0
	fmt.Println(threeSumClosest([]int{-1,2,1,-4}, 2))	//2
	fmt.Println(threeSumClosest([]int{-1,2,1,-4}, 1))	//2
	fmt.Println(threeSumClosest([]int{1,1,1,0}, -100))	//2
}
