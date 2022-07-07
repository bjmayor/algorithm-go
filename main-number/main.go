package main

import "fmt"

func main_number(nums []int) int {
	var left = 0;
	var right = len(nums)-1;
	for left < right {
		// 2数不相等，移除，这里是都放到数组的左边。 left 左边都是被移除的，右边是剩下的。
		if nums[left]!=nums[right] {
			// right 和 left+1 交换。
			temp := nums[right]
			nums[right] = nums[left+1]
			nums[left+1] = temp
			left += 2;
			if left >= right {
				if left+1 < len(nums) {
					return nums[left+1]
				} else {
					return 0
				}
			}
		} else {
			right--;
			if left>=right {
				return nums[left];
			}
		}
	}
	return 0
}

func main()  {
	cases := []struct {
		nums []int
		expected int
	}{
		//{
		//	[]int{
		//		1, 9, 2, 8, 1, 3, 2, 1, 1, 1, 1,
		//	}, 1,
		//},
		//{
		//	[]int{
		//		1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 5,
		//	}, 1,
		//},
		//{
		//	[]int{
		//		3, 9, 2, 8, 1, 3, 1, 1, 1, 1, 1,
		//	}, 1,
		//},
		//{
		//	[]int{
		//		1, 1, 1, 1, 1, 3, 2, 1, 1, 1, 1,
		//	}, 1,
		//},
		{
			[]int{
				1, 9, 2, 8, 1, 3, 2, 1, 2, 3, 5,
			}, 0,
		},
	}
	for i,t := range cases {
		real := main_number(t.nums)
		if (real!=t.expected) {
			fmt.Sprintf("%d error, expected %, real is %d",i, t.expected, real )
		}
	}
}
