package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 控制成 前短， 后长
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	// 假设 i,j 即为中位数（需要满足i左边+j左边的 个数和 刚好=half,
	// 探测，缩小范围 满足时会有nums1[i-1]<=nums2[j] nums2[j]<=nums1[i]
	m := len(nums1)
	n := len(nums2)
	imin, imax, half := 0, m, (m+n+1)/2

	var mid1, mid2 int
	for imin <= imax { //i 还有可选值
		i :=  (imin+imax)/2
		j := half-i
		a1 := math.MinInt32
		a2 := math.MaxInt32
		if i > 0 {
			a1 = nums1[i-1]
		}
		if i < m {
			a2 = nums1[i]
		}

		b1 := math.MinInt32
		b2 := math.MaxInt32
		if j > 0 {
			b1 = nums2[j-1]
		}
		if j < n {
			b2 = nums2[j]
		}

		if a1 <= b2 {
			mid1 = max(a1,b1)
			mid2 = min(a2,b2)
			imin++
		} else {
			imax--
		}
	}
	if (m+n) % 2 == 1 {
		return float64(mid1)
	}
	return float64(mid1+mid2)/2
}

func max(a,b int) int  {
	if a > b {
		return a
	}
	return b
}

func min(a,b int) int  {
	if a > b {
		return b
	}
	return a
}


func main() {
	tests := []struct {
		Nums1    []int
		Nums2    []int
		Expected float64
	}{
		//{[]int{1, 3}, []int{2}, 2.0},
		//{[]int{1, 2}, []int{3, 4}, 2.5},
		//{[]int{1, 2,5}, []int{3, 4}, 3.0},
		//{[]int{1, 2,5}, []int{3, 4,6}, 3.5},
		//{[]int{0, 0}, []int{0, 0}, 0},
		{[]int{}, []int{2, 3}, 2.5},
		{[]int{3}, []int{-2, -1}, -1.0},
		{[]int{1,2}, []int{-1, 3}, 1.5},
		{[]int{1,2,2}, []int{1,2, 3}, 2.0},
		{[]int{1}, []int{2,3, 4}, 2.5},
	}
	for _, t := range tests {
		var real = findMedianSortedArrays(t.Nums1, t.Nums2)
		if real != t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, "data:", t.Nums1, t.Nums2)
		}
	}
}
