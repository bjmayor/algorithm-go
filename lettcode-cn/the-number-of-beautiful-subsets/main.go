package main

import (
	"fmt"
	"sort"
)

/**
 * @description 计算美丽子集的数量
 * @param nums 整数数组
 * @param k 给定的差值
 * @return int 美丽子集的数量
 *
 * 算法思想:
 * 1. 首先对数组进行排序
 * 2. 使用回溯法遍历所有可能的子集
 * 3. 空集不算美丽子集，所以最后结果减1
 */
func beautifulSubsets(nums []int, k int) int {
	sort.Ints(nums) // 先排序

	// 使用哈希表记录已选择的数字
	selected := make(map[int]bool)

	/**
	 * 回溯算法的通用模板:
	 * 1. 定义结束条件
	 * 2. 做出选择
	 * 3. 递归调用
	 * 4. 撤销选择(回溯)
	 *
	 * @param index 当前处理的元素索引
	 * @return 符合条件的子集数量
	 */
	var backtrack func(int) int
	backtrack = func(index int) int {
		// 1. 结束条件: 已处理完所有数字
		if index == len(nums) {
			// 如果选择了至少一个数字，则为有效子集
			if len(selected) > 0 {
				return 1
			}
			return 0
		}

		// 2a. 选择一: 不选当前数字
		count := backtrack(index + 1)

		// 2b. 选择二: 尝试选择当前数字
		// 检查是否可以选择当前数字 - 排序后只需检查nums[index]-k
		canSelect := true
		if _, exists := selected[nums[index]-k]; exists {
			canSelect = false
		}

		// 如果可以选择当前数字
		if canSelect {
			// 3. 做出选择
			selected[nums[index]] = true

			// 4. 递归调用
			count += backtrack(index + 1)

			// 5. 撤销选择(回溯)
			delete(selected, nums[index])
		}

		return count
	}

	return backtrack(0)
}

/**
 * 回溯算法通用模板:
 *
 * func backtrack(参数) {
 *     // 1. 结束条件
 *     if 满足结束条件 {
 *         保存结果
 *         return
 *     }
 *
 *     // 2. 遍历所有可能的选择
 *     for 选择 in 选择列表 {
 *         // 3. 做出选择
 *         进行选择
 *
 *         // 4. 递归调用
 *         backtrack(新参数)
 *
 *         // 5. 撤销选择(回溯)
 *         撤销选择
 *     }
 * }
 */

func main() {
	nums := []int{1, 2, 3, 4}
	k := 1
	fmt.Println(beautifulSubsets(nums, k)) // 期望输出: 7
}
