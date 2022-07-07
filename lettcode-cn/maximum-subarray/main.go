package main

import "math"

// 1. 暴力枚举
func baoli(nums []int) int {
	largest := math.MinInt64
	for i := 0; i < len(nums); i++ {
		result := nums[i]
		if result > largest {
			largest = result
		}
		for j := i + 1; j < len(nums); j++ {
			result += nums[j]
			if result > largest {
				largest = result
			}
		}
	}
	return largest
}

//2 动态规划
func dp(nums []int) int {
	// 包含i的累计和
	maxall := map[int]int{
		0:nums[0],
	}
	for i, v := range nums[1:] {
		maxall[i+1]	= v
		if maxall[i] > 0 {
			maxall[i+1] += maxall[i]
		}
	}
	largest := nums[0]
	for _,v := range maxall {
		if v > largest {
			largest = v
		}
	}
	return largest
}


/***
3. 如果当前累计和>0, 那么可以扩展。
否则 重新开始。
 */
func a(nums []int) int  {
	res := nums[0];
	sum := 0;
	for _, num := range nums {
		if sum > 0 {
			sum += num;
		} else {
			sum = num;
		}
		if sum > res {
			res = sum
		}
	}
	return res;
}

// 合并才关键
func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum + r.lSum)
	rSum := max(r.rSum, r.iSum + l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

// 4. 分治法
func get(nums []int, l, r int) Status {
	if (l == r) {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m + 1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Status struct {
	lSum, // 包含l的最大累计和
	rSum, // 包含r的最大累计和
	mSum,  // 最大连续子数组的累计和
	iSum int // 子数组的累计和
}


func maxSubArray(nums []int) int {
	return get(nums, 0, len(nums) - 1).mSum
}

func main() {
	println(maxSubArray([]int{8,-19,5,-4,20}))                            //21
	println(maxSubArray([]int{-2,-1}))                            //-1
	println(maxSubArray([]int{-9}))                            //-9
	println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) //6
	println(maxSubArray([]int{1}))                             //1
	println(maxSubArray([]int{5, 4, -1, 7, 8}))                //23
}
