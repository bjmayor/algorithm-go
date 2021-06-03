package main

import "fmt"

func main()  {
	for i, t := range tests {
		var real = findMaxLength(t.Nums)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real, "\t case",t.Nums)
		}
	}
}

// 暴力求解
func findMaxLength1(nums []int) int {
	var result int
	for i:=0;i<len(nums)-1;i++{
		sum :=0
		for j:=i;j<len(nums);j++ {
			sum+=nums[j]
			n := j-i+1
			if sum*2 == n && n >result {
				result = n
			}
		}
	}
	return result
}

// 前缀和
func findMaxLength(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	for i:=0;i<len(nums);i++{
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	result := 0
	pre := 0
	set := map[int]int{ 0:-1}
	for i:=0;i<len(nums);i++{
		 pre = pre+nums[i]
		if j,ok := set[pre];ok {
			size := i-j
			if size>result  {
				result = size
			}
		} else {
			set[pre] = i
		}
	}
	return result
}

// dp求解
func findMaxLength2(nums []int) int {
	var result int
	if len(nums) < 2 {
		return 0
	}
	dp := make([]int, len(nums), len(nums))
	for i:=1;i<len(nums);i++ {
		// 上一个元素有解, 尝试连接已有解。
		var connectedPre int
		if dp[i-1] > 0 {
			pre := i-dp[i-1]-1
			if pre >=0  && nums[pre]+nums[i]==1{
				dp[i] = dp[i-1]+2
				for {
					pre := i-dp[i]
					// 链接更早的解
					if pre>=0 && dp[pre]>0 {
						dp[i] += dp[pre]
					}	else {
						break
					}
				}
				connectedPre = dp[i]
				if dp[i] > result {
					result = dp[i]
				}
			}
		}
		// 直接和上一个元素连接试试
		if nums[i-1]+nums[i] == 1{
			dp[i] = 2
			for {
				pre := i-dp[i]
				// 如果dp[i-2]有解，需要连接上
				if pre>=0 && dp[pre]>0 {
					dp[i] += dp[pre]
				}	else {
					break
				}
			}
			if dp[i] < connectedPre {
				dp[i] = connectedPre
			}

			if dp[i] > result {
				result = dp[i]
			}
		}
	}
	return result
}