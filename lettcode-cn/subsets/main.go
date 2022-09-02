package main

import "fmt"

func subsets(nums []int) [][]int {
	var ans [][]int
	ans = append(ans, []int{})
	for i := 0; i < len(nums); i++ {
		size := len(ans)
		for j := 0; j < size; j++ {
			n := make([]int, len(ans[j])+1)
			copy(n, ans[j])
			n[len(ans[j])] = nums[i]
			ans = append(ans, n)
		}
	}
	return ans
}

func main() {
	//fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}
