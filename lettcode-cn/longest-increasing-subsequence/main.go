package main

import "fmt"
func lengthOfLIS(nums []int) int {
	return lengthOfLIS2(nums)
}
// dp求解 n平方算法
func lengthOfLIS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var result  = 1
	var dp  = make([]int, len(nums), len(nums)) //最后一个元素为i的最长递增序列长度
	dp[0 ] =1
	for i:=1;i<len(nums);i++ {
		max := 0
		dp[i] = 1
		for j:=i;j>=0;j-- {
			if nums[j] < nums[i]  && dp[j]>max {
				dp[i] = dp[j]+1
				max = dp[j]
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return  result
}

//  二分查找 nlgn 算法
func lengthOfLIS2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var lts  = make([]int, 0)
	lts = append(lts, nums[0])
	for i:=1;i<len(nums);i++ {
		if nums[i] > lts[len(lts)-1] {
			lts = append(lts, nums[i])
		} else {
			left := 0
			right := len(lts)-1
			//找第一个>=num[i]的数
			for left < right {
				mid := (left+right)/2
				if lts[mid] < nums[i] {
					left = mid+1
				} else {
					right = mid
				}
			}
			lts[left] = nums[i]

		}
	}
	return len(lts)
}

func main()  {
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))//4
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,4}))//3
	fmt.Println(lengthOfLIS([]int{2,2}))//1
	fmt.Println(lengthOfLIS([]int{4,10,4,3,8,9}))//3
	fmt.Println(lengthOfLIS([]int{3,5,6,2,5,4,19,5,6,7,12}))//6
}
