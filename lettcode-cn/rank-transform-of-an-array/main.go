package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	m := make(map[int]int)
	arr2 := make([]int, len(arr), len(arr))
	copy(arr2, arr)
	sort.Ints(arr2)
	rank := 1
	pre := arr2[0]
	m[pre] = rank
	for i:=1;i<len(arr2);i++{
		if arr2[i]	> pre {
			rank++
		}
		pre = arr2[i]
		m[pre] = rank
	}

	for i, v := range arr {
		arr[i] = m[v]
	}
	return arr
}

func main()  {
	fmt.Println(arrayRankTransform([]int{40,10,20,30}))
}
