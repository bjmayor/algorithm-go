package main

import (
	"fmt"
	"math"
)

/*
* jump 计算跳到数组最后一个位置所需的最少跳跃次数
* 算法思想:
* 1. 使用动态规划,dp[i]表示跳到位置i所需的最少跳跃次数
* 2. 状态转移:
*    - 对于位置i,检查之前的每个位置j
*    - 如果位置j可达(dp[j]!=math.MaxInt32)且从j能跳到i(nums[j]>=i-j)
*      则更新dp[i]为dp[j]+1
* 3. 初始化:
*    - dp[0] = 0 表示起点所需跳跃次数为0
* 4. 返回dp[len(nums)-1]即为所需最少跳跃次数
*
* 时间复杂度: O(n^2) - 两层循环遍历
* 空间复杂度: O(n) - dp数组
 */
func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < i; j++ {
			if nums[j] >= i-j {
				dp[i] = min(dp[i], dp[j]+1)
				break
			}
		}
	}
	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // 2
	fmt.Println(jump([]int{2, 3, 0, 1, 4})) // 2
}
