package main

import (
	"fmt"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	m := make(map[int][]int)
	data := make([]int,0)
	for i:=0;i<len(mat);i++ {
		cnt :=0
		for j:=0;j<len(mat[i]);j++ {
			cnt+=mat[i][j]
		}
		v := m[cnt]
		v = append(v, i)
		m[cnt] = v
	}
	for k,_ := range m {
		data = append(data,k)
	}
	sort.Ints(data)
	var res []int
	for i:=0;i<k;i++ {
		res = append(res, m[data[i]]...)
		if len(res) >=k {
			return res[:k]
		}
	}
	return res
}

func kWeakestRows2(mat [][]int, k int) []int {
	data := make([]int,0)
	for i:=0;i<len(mat);i++ {
		cnt :=0
		for j:=0;j<len(mat[i]);j++ {
			cnt+=mat[i][j]
		}
		data = append(data, cnt*1000+i)
	}
	sort.Ints(data)
	var res []int
	for i:=0;i<k;i++ {
		res = append(res, data[i]%1000)
	}
	return res
}


func main()  {
	//fmt.Println(kWeakestRows([][]int{
	//{1,1,0,0,0},
	//	{1,1,1,1,0},
	//	{1,0,0,0,0},
	//	{1,1,0,0,0},
	//	{1,1,1,1,1},
	//},3))

	fmt.Println(kWeakestRows2([][]int{
		{1,0},
		{1,0},
		{1,0},
		{1,1},
	},4))
}
