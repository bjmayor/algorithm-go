package main

import (
	"fmt"
	"sort"
)

/**
 * 计算每个查询的最大美丽值
 *
 * 优化算法思想:
 * 1. 对物品按价格排序
 * 2. 更新每个价格点的最大美丽值
 * 3. 对每个查询使用二分查找找到对应价格位置
 *
 * 时间复杂度：O(NlogN + QlogN)，其中N是物品数量，Q是查询数量
 * 空间复杂度：O(N)
 *
 * @param items [][]int - 物品数组，每个物品由价格和美丽值组成
 * @param queries []int - 查询数组
 * @return []int - 每个查询对应的最大美丽值数组
 */
func maximumBeauty(items [][]int, queries []int) []int {
	n := len(items)
	if n == 0 {
		return make([]int, len(queries))
	}

	// 1. 按价格排序
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	// 2. 更新每个价格点的最大美丽值
	for i := 1; i < n; i++ {
		if items[i][1] < items[i-1][1] {
			// 相同价格，保留最大美丽值
			items[i][1] = items[i-1][1]
		}
	}

	ans := make([]int, len(queries))
	// 对每个查询使用二分查找
	for i, query := range queries {
		if query < items[0][0] {
			ans[i] = 0
			continue
		}
		pos := sort.Search(n, func(j int) bool {
			return items[j][0] > query
		}) - 1
		if pos >= 0 {
			ans[i] = items[pos][1]
		}
	}
	return ans
}

func main() {
	items := [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}
	queries := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(maximumBeauty(items, queries)) // [2, 4, 5, 5, 6, 6]
}
