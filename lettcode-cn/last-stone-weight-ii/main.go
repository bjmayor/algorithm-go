package main

import (
	"sort"
	"strconv"
)

var set = make(map[string]int)
// 暴力求解
func lastStoneWeightII1(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	key := hash(stones)
	if v, ok := set[key];ok {
		return v
	}
	var result = 101
	for i:=0;i<len(stones);i++ {
		for j:=i+1;j<len(stones);j++ {
			left := removeIndex(stones, map[int]struct{}{
				i: {},
				j:{},
			})
			if stones[i] > stones[j] {
				left = append(left, stones[i]-stones[j])
			} else if stones[i] < stones[j]{
				left = append(left, stones[j]-stones[i])
			}
			v := lastStoneWeightII(left)
			if v < result {
				result = v
			}
		}
	}
	set[key] = result
	return result
}

func hash(data []int) string  {
	sort.Ints(data)
	result :=""
	for i:=0;i<len(data);i++ {
		result += ","+strconv.Itoa(data[i])
	}
	return result
}

func removeIndex(stones []int, idx map[int]struct{}) []int  {
	var result []int
	for i:=0;i<len(stones);i++ {
		if _, ok := idx[i];ok {
			continue
		}
		result = append(result, stones[i])
	}
	return result
}

func lastStoneWeightII(stones []int) int {
	for i:=0;i<101;i++ {
		if findTargetSumWays(stones,i) > 0 {
			return i
		}
	}
	return 101
}
func findTargetSumWays(nums []int, target int) int {
	var total int
	for i:=0;i<len(nums);i++ {
		total += nums[i]
	}
	if (total+target) % 2 ==1 || total < target {
		return 0
	}
	return findTargetSumPositiveWays(nums, (total+target)/2)
}
func findTargetSumPositiveWays(nums []int ,target int) int {
	var dp = make([]int, target+1, target+1)
	dp[0] = 1
	for i:=0;i<len(nums);i++ {
		for j:=target;j>=nums[i];j-- {
			dp[j] = dp[j]+dp[j-nums[i]]
		}
	}
	return dp[target]
}

func main()  {
	//println(lastStoneWeightII([]int{2,7,4,1,8,1}))
	tests := []struct {
		Nums []int
		Expected int
	}{
		{ []int{2,7,4,1,8,1},1},
		{ []int{31,26,33,21,40},5},
		{ []int{1,2},1},
		//{ []int{1,1,2,3,5,8,13,21,34,55,89,14,23,37,61,98},1},

	}
	for i,item:= range tests {
		real := lastStoneWeightII(item.Nums)
		if real != item.Expected {
			println(i," failed")
		}
	}
}