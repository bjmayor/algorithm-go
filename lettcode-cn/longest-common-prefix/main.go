package main

import (
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) ==0 {
		return ""
	}
	min := math.MaxInt32
	idx:=0
	for i, str := range strs {
		if len(str)	< min {
			min = len(str)
			idx  = i
		}
	}
	for i:=0;i<min;i++ {
		c := strs[0][i]
		for j:=1;j<len(strs);j++ {
			if strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[idx]
}

