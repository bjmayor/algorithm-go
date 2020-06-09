package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := make([][]int,0)
	sort.Ints(nums)
	for i:=0; i<len(nums)-2; i++{

		v := nums[i]
		y :=i+1
		z := len(nums)-1
		left := 0-v
		for y<z {
			sum :=  nums[y]	+ nums[z]
			if sum < left {
				y++
			} else if sum > left {
				z--
			} else {
				result = append(result, []int{v,nums[y],nums[z]})
				for y+1 < z {
					if nums[y+1] == nums[y] {
						y= y+1
					} else {
						break
					}
				}
				y++
				for y+1 < z {
					if nums[z-1] == nums[z] {
						z--
					} else {
						break
					}
				}
				z--
			}
		}
		for i+1 < len(nums)-2 && nums[i+1] == nums[i]{
			i = i+1
		}
	}
	return result

}

func main()  {
	fmt.Println(threeSum([]int{
		-1,0,1,2,-1,-4,
	}))
}

