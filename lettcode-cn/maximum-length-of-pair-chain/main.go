package main

import (
	"math"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] > pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})
	var ans int
	var cur = math.MinInt32
	for _, pair := range pairs {
		if pair[0] > cur {
			cur = pair[1]
			ans++
		}
	}
	return ans
}

func main() {
	println(findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
}
