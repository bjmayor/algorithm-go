package main

import (
	"fmt"
)
//线性搜索
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	m := make(map[int]int)
	for i:=0;i<len(nums);i++ {
		id := getId(nums[i], t)
		if _, ok := m[id]; ok {
			return true
		} else {
			if v, ok := m[id+1]; ok && v-nums[i] <=t {
				return true
			}
			if v, ok := m[id-1]; ok && nums[i]-v <=t {
				return true
			}
			m[id] = nums[i]
			if i+1 > k  && i-k>=0{
				pre := getId(nums[i-k], t)
				delete(m, pre)
			}
		}
	}

	return false
}


//0~t 为0，-t+1~-1 为-1
func getId(v, t int) int  {
	if v >=0 {
		return v/(t+1)
	} else {
		return (v+1)/(t+1)-1
	}
}



func main()  {
	tests := []struct{
		Nums []int
		K int
		T int
		Expected bool
	} {
		//{ []int{1,2,3,1},3,0,true},
		//{ []int{1,0,1,1},1,2,true},
		//{ []int{1,5,9,1,5,9},2,3,false},
		{ []int{-3,3},2,4,false},
	}
	for i, t := range tests {
		var real = containsNearbyAlmostDuplicate(t.Nums, t.K, t.T)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
