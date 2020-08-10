package main

import "fmt"

func countBinarySubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	var ret int
	var nums []int
	var i int
	for i<len(s) {
		c := s[i]
		n := 0
		for i<len(s) && s[i] == c {
			n++
			i++
		}
		nums = append(nums, n)
	}
	for i:=0;i<len(nums)-1;i++ {
		ret += min(nums[i], nums[i+1])
	}
	return ret
}

func min(x,y int) int  {
	 if x < y {
	 	return x
	 }
	 return y
}

func main()  {
	fmt.Println(countBinarySubstrings("00110011"))
}
