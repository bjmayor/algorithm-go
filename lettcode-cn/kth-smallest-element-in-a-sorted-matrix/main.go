package main

import (
	"fmt"
)
type Coor struct {
	i int
	j int
}
//最简单直接的做法，效率低
//func kthSmallest(matrix [][]int, k int) int {
//	var res =matrix[0]
//	for i:=1;i<len(matrix);i++ {
//		res = merge(res, matrix[i])
//	}
//	return res[k-1]
//}

func merge(a,b []int) []int {
	i:=0
	j:=0
	k:=0
	var res = make([]int, len(a)+len(b))
	for i<len(a) && j<len(b) {
		if a[i]	< b[j] {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
		k++
	}

	if i==len(a) {
		for  j<len(b) {
			res[k] = b[j]
			j++
			k++
		}
	}
	if j == len(b) {
		for  i<len(a) {
			res[k] = a[i]
			i++
			k++
		}
	}
	return res
}

//二分查找
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right - left) / 2
		if check(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

//<=mid 的 个数是否 >=k, 如果是，说明在左上角，否则在右下角
func check(matrix [][]int, mid, k, n int) bool {
	i, j := n - 1, 0
	num := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			num += i + 1
			j++
		} else {
			i--
		}
	}
	return num >= k
}


func main()  {
	fmt.Println(kthSmallest([][]int {
		{0,  5,  9},
		{10, 21, 33},
		{22, 33, 55},
	},8)) //13
}