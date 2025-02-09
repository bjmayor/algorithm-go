package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	var dfs func(index int, path []int)
	dfs = func(index int, path []int) {
		// 添加当前路径, 不选择index
		result = append(result, append([]int{}, path...))
		for i := index; i < len(nums); i++ {
			// 跳过重复的元素
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			// 选择index
			dfs(i+1, append(path, nums[i]))
		}
	}
	dfs(0, []int{})
	return result
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
