package main

import (
	"fmt"
)

/*
2012. 数组美丽值求和

给你一个下标从 0 开始、长度为 n 的整数数组 nums 。

对于每个下标 i（1 <= i <= nums.length - 2），nums[i] 的 美丽值 等于：

2，对于所有 0 <= j < i 且 i < k <= nums.length - 1 ，满足 nums[j] < nums[i] < nums[k]
1，如果满足 nums[i - 1] < nums[i] < nums[i + 1] ，且不满足前面的条件
0，如果上述条件全部不满足
返回符合 1 <= i <= nums.length - 2 的所有 nums[i] 的 美丽值的总和
*/

func sumOfBeauties(nums []int) int {
	n := len(nums)
	// 预处理左侧最大值
	leftMax := make([]int, n)
	leftMax[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], nums[i])
	}

	// 预处理右侧最小值
	rightMin := make([]int, n)
	rightMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMin[i] = min(rightMin[i+1], nums[i])
	}

	ans := 0
	// 计算美丽值
	for i := 1; i < n-1; i++ {
		if leftMax[i-1] < nums[i] && nums[i] < rightMin[i+1] {
			ans += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			ans++
		}
	}
	return ans
}

// 辅助函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(sumOfBeauties([]int{1, 2, 3}))    // 2
	fmt.Println(sumOfBeauties([]int{2, 4, 6, 4})) // 1
	fmt.Println(sumOfBeauties([]int{3, 2, 1}))    // 0
}
