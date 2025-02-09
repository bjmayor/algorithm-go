package main

import (
	"fmt"
	"sort"
)

// 对于这类寻找所有可行解的题，我们都可以尝试用「搜索回溯」的方法来解决。
/*
* combinationSum 计算数组中所有和为 target 的组合
* @param candidates 候选数字数组
* @param target 目标和
* @return [][]int 所有可能的组合
*
* 解题思路:
* 1. 先对数组排序,方便剪枝
* 2. 使用回溯(DFS)搜索所有可能组合:
*    - 从当前位置开始遍历,避免重复组合
*    - 当前数字可以重复使用,所以下次搜索的起始位置仍为当前位置i
*    - 如果当前数字大于target,后面更大的数字都不需要尝试
* 3. 记录路径,当target减到0时得到一个解
 */

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 排序方便剪枝
	res = [][]int{}
	dfs(candidates, target, 0, []int{})
	return res
}

/*
* dfs 深度优先搜索所有可能组合
* @param candidates 候选数字数组
* @param target 当前要凑成的目标和
* @param index 当前搜索的起始位置
* @param path 当前路径
 */
func dfs(candidates []int, target int, index int, path []int) {
	if target == 0 { // 找到一个解
		res = append(res, append([]int{}, path...))
		return
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target { // 剪枝:当前数字太大,后面更大的数字都不用试了
			break
		}
		dfs(candidates, target-candidates[i], i, append(path, candidates[i]))
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
