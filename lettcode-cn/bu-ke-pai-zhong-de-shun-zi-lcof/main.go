package main

import "fmt"

func isStraight(nums []int) bool {
	data := make([]int, 14)
	for _, v := range nums {
		data[v]++
	}
	for i, v := range data {
		if i!=0 && v==1 {
			sq :=1
			for j:=i+1;j<14 && j<i+5;j++ {
				sq++
				if data[j] <= 1 {
					if data[j] == 0 {
						if data[0] > 0 {
							data[0]--
						}else {
							return false
						}
					}
				} else {
					return false
				}
			}
			if sq == 5 || sq>=3 && sq+data[0] == 5 {
				return true
			}
			return false
		}
	}
	return false
}

func main()  {
	tests := []struct{
		Nums []int
		Expected bool
	} {
		{ []int{1,2,3,4,5},true},
		{ []int{0,0,1,2,5},true},
		{ []int{0,0,12,13,11},true},
	}
	for i, t := range tests {
		var real = isStraight(t.Nums)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}