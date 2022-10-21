package main

import "sort"

func lastStoneWeight(stones []int) int {
	sort.Slice(stones, func(i, j int) bool {
		return stones[i] > stones[j]
	})
	for len(stones) > 1 {
		left := stones[0] - stones[1]
		stones = stones[2:]
		if left > 0 {
			idx := sort.Search(len(stones), func(i int) bool {
				return stones[i] < left
			})
			if idx == len(stones) {
				stones = append(stones, left)
			} else {
				stones = append(stones[:idx], append([]int{left}, stones[idx:]...)...)
			}
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

func main() {
	println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}
