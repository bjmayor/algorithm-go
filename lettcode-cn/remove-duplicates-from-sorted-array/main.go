package main

import "fmt"

func removeDuplicates(nums []int) int {
	j := 1
	for i:=1;i<len(nums);i++ {
		if nums[i] != nums[i-1]{
			if i!=j {
				nums[j] = nums[i]
			}
			j++
		}
	}
	return j
}



func main()  {
	nums := []int{
		-1,-1,0,2,3,
	}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}