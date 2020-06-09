package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := make([][]int,0)
	sort.Ints(nums)
	for i:=0; i<len(nums)-3; i++{
		for j:=i+1;j<len(nums)-2;j++ {
			y :=j+1
			z := len(nums)-1
			left := target-nums[i]-nums[j]
			for y<z {
				sum :=  nums[y]	+ nums[z]
				if sum < left {
					y++
				} else if sum > left {
					z--
				} else {
					result = append(result, []int{nums[i],nums[j],nums[y],nums[z]})
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
			for j+1 < len(nums)-2 && nums[j+1] == nums[j]{
				j = j+1
			}
		}
		for i+1 < len(nums)-3 && nums[i+1] == nums[i]{
			i = i+1
		}

	}
	return result
}


func main()  {
	fmt.Println(fourSum([]int{
		1, 0, -1, 0, -2, 2,
	},0))
}
