package main

import (
	"sort"
)

func sumSubseqWidths(nums []int) (ans int) {
	const mod int = 1e9 + 7
	sort.Ints(nums)
	res, x, y := 0, nums[0], 2
	for _, num := range nums[1:] {
		res = (res + num*(y-1) - x) % mod
		x = (x*2 + num) % mod
		y = y * 2 % mod
	}
	return (res + mod) % mod
}

func main() {
	println(sumSubseqWidths([]int{2, 1, 3}))    //6
	println(sumSubseqWidths([]int{3, 7, 2, 3})) //35
}
