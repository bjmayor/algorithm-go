package main

import "fmt"

/*
* canJump 判断是否能跳到数组最后一个位置
* 优化算法思想:
* 1. 使用贪心算法,维护一个变量maxReach表示当前能到达的最远位置
* 2. 遍历数组,更新maxReach
* 3. 如果maxReach达到或超过最后一个位置,返回true
* 4. 如果遍历结束后maxReach仍然小于最后一个位置,返回false
*
* 时间复杂度: O(n) - 单次遍历
* 空间复杂度: O(1) - 只使用常数空间
 */
func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			// 如果当前位置超出最大可达范围, 返回false
			return false
		}
		// 更新最大可达范围
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= len(nums)-1 {
			// 如果已经可以到达最后一个位置, 返回true
			return true
		}
	}
	return false
}

// 辅助函数: 返回较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false
}
