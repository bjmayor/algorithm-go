package main

import "fmt"

func nextPermutation(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	j := len(nums)-1
	for j>=1 {
		//比前面的小，则找个刚好比前面大的，交换，剩下的升序
		if nums[j] > nums[j-1] {
			for i:=len(nums)-1;i>=j;i-- {
				if nums[i] > nums[j-1] { //找到第一个比j大的，交换，剩下的数字从降序改为升序
					nums[i],nums[j-1] = nums[j-1], nums[i]
					asc(nums, j, len(nums)-1)
					return
				}
			}
			return
		} else {
			j--
			if j==0 {
				asc(nums, 0, len(nums)-1)
				return
			}
		}
	}
}
func asc(nums []int, start, end int)  {
	for k:=start;k<=(start+end)/2;k++{
		nums[k],nums[end-k+start] = nums[end-k+start], nums[k]
	}
}

func main()  {
	nums := []int{1,2,3}
	//nextPermutation(nums)
	//fmt.Println(nums) //1,3,2

	nums = []int{1,3,2}
	nextPermutation(nums)
	fmt.Println(nums) //2,1,3

	//nums = []int{3,2,1}
	//nextPermutation(nums)
	//fmt.Println(nums) //1,2,3
	//
	//nums = []int{1,1,5}
	//nextPermutation(nums)
	//fmt.Println(nums)//1,5,1
}