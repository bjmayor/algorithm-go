package main

import "sort"

func trimMean(arr []int) float64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	start := len(arr) / 20
	end := len(arr) * 19 / 20
	part := arr[start:end]
	sum := 0
	for _, v := range part {
		sum += v
	}
	return float64(sum) / float64(len(part))
}
