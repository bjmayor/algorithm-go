package maximizehappinessofselectedchildren

import "slices"

func maximumHappinessSum(happiness []int, k int) int64 {
	// 降序排序，优先选择幸福值最大的孩子
	// 因为所有未选中的孩子幸福值同时减少，所以优先选大的一定最优
	slices.Sort(happiness)
	slices.Reverse(happiness)
	
	var total int64
	// 第i个被选中的孩子，经历了i轮减少，幸福值为 happiness[i] - i
	for i := 0; i < k; i++ {
		// 幸福值不能为负数，如果减少后小于0则贡献0
		total += max(int64(happiness[i]-i), 0)
	}
	return total
}
