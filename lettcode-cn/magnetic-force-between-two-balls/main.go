package main

import (
	"fmt"
	"sort"
)

/**
 * maxDistance 函数用于计算在给定位置数组中放置 m 个球时，球之间的最大最小距离。
 * 算法思想：使用二分查找法来确定最大最小距离。首先对位置数组进行排序，
 * 然后在可能的距离范围内进行二分查找。通过检查在当前距离下能否放置 m 个球，
 * 来调整二分查找的上下界。
 *
 * @param position 整数数组，表示可放置球的位置
 * @param m 整数，表示需要放置的球的数量
 * @return int 返回球之间的最大最小距离
 */
func maxDistance(position []int, m int) int {
	// 对位置数组进行排序
	sort.Ints(position)
	// 初始化二分查找的范围
	left, right := 1, position[len(position)-1]-position[0]

	// 定义一个函数来检查在给定距离下能否放置 m 个球
	canPlaceBalls := func(distance int) bool {
		count, lastPosition := 1, position[0]
		for i := 1; i < len(position); i++ {
			if position[i]-lastPosition >= distance {
				count++
				lastPosition = position[i]
				if count >= m {
					return true
				}
			}
		}
		return false
	}

	// 二分查找最大最小距离
	for left <= right {
		mid := (left + right) / 2
		if canPlaceBalls(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func main() {
	// 调用maxDistance函数并输出结果
	fmt.Println("最大距离:", maxDistance([]int{1, 2, 3, 4, 7}, 3))                 // 3
	fmt.Println("最大距离:", maxDistance([]int{5, 4, 3, 2, 1, 1000000000}, 2))     //999999999
	fmt.Println("最大距离:", maxDistance([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4)) // 3
}
