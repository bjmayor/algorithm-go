package main

import (
	"fmt"
	"sort"
)

/**
 * 2279. 装满花园的最大总花数
 *
 * 算法思想:
 * 1. 贪心策略
 *    - 考虑将前i个花园填满到target
 *    - 对剩余花园，尽可能提高其最小值
 * 2. 使用前缀和和二分查找优化:
 *    - 预计算前缀和，避免重复计算
 *    - 使用二分查找快速定位分界点
 *    - 优化遍历过程，减少重复计算
 *
 * @param flowers []int - 每个花园的初始花数量
 * @param newFlowers int64 - 可以额外种植的花的数量
 * @param target int - 完整花园所需的花数量
 * @param full int - 完整花园的得分
 * @param partial int - 不完整花园的得分系数
 * @return int64 - 最大可能的总分
 */
func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	n := len(flowers)
	remain := int(newFlowers) - n*target

	// 预处理：将超过target的花园处理为target
	for i := 0; i < n; i++ {
		flowers[i] = min(flowers[i], target)
		remain += flowers[i]
	}

	// 快速判断边界情况
	if remain == int(newFlowers) {
		return int64(n * full)
	}
	if remain >= 0 {
		return int64(max(n*full, (n-1)*full+(target-1)*partial))
	}

	sort.Ints(flowers)

	// 计算前缀和数组
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + flowers[i]
	}

	ans := 0
	j := 0
	// 枚举完整花园的数量
	for i := 1; i <= n; i++ {
		remain += (target - flowers[i-1])
		if remain < 0 {
			continue
		}

		// 二分查找最大可能的j
		newJ := sort.Search(i+1, func(mid int) bool {
			if mid == 0 {
				return false
			}
			return flowers[mid-1]*mid > remain+preSum[mid]
		})
		if newJ > 0 {
			newJ--
		}
		j = max(j, newJ)

		// 计算平均值
		if j > 0 {
			avg := (remain + preSum[j]) / j
			ans = max(ans, (n-i)*full+avg*partial)
		}
	}

	return int64(ans)
}

// 辅助函数：返回两个int中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 辅助函数：返回两个int中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 测试用例1
	flowers1 := []int{1, 3, 1, 1}
	newFlowers1 := int64(7)
	target1 := 6
	full1 := 12
	partial1 := 1
	fmt.Println(maximumBeauty(flowers1, newFlowers1, target1, full1, partial1)) //14

	// 测试用例2
	flowers2 := []int{2, 4, 5, 3}
	newFlowers2 := int64(10)
	target2 := 5
	full2 := 2
	partial2 := 6
	fmt.Println(maximumBeauty(flowers2, newFlowers2, target2, full2, partial2)) //30

	// 测试用例3
	flowers3 := []int{8, 20, 13}
	newFlowers3 := int64(12)
	target3 := 12
	full3 := 4
	partial3 := 3
	fmt.Println(maximumBeauty(flowers3, newFlowers3, target3, full3, partial3)) //41

	// 测试用例4
	flowers4 := []int{89287, 5538, 37141, 72522, 84502, 44451, 24432, 2324, 72779, 57060, 99178, 6, 29440, 53664, 76197, 46742, 17384, 22072, 33067, 66274, 19323, 72943, 12914, 91475, 96826, 84847, 28100, 89590, 34977, 74052, 4813, 24563, 97491, 71687, 8533, 49262, 2265, 10553, 63902, 19647, 27006, 64548, 89892, 64046, 72766, 34623, 17297, 21417, 70630, 93469, 83379, 19483, 93842, 65968, 28401, 1889, 24441, 99401, 37907, 13794, 3640, 95432, 36875, 10200, 95360, 10829, 96763, 15900, 8490, 68972, 52537, 72458, 95269}
	newFlowers4 := int64(42)
	target4 := 4534
	full4 := 32415
	partial4 := 11040
	fmt.Println(maximumBeauty(flowers4, newFlowers4, target4, full4, partial4)) //2734140
}
