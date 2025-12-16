package main

type result struct {
	dp0  []int // 父节点不购买时，各预算下的最大收益
	dp1  []int // 父节点购买时（当前节点可享受半价），各预算下的最大收益
	size int   // 子树最大可能开销（用于剪枝）
}

func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	// 构建邻接表
	g := make([][]int, n)
	for _, edge := range hierarchy {
		parent, child := edge[0]-1, edge[1]-1
		g[parent] = append(g[parent], child)
	}

	var dfs func(u int) result
	dfs = func(u int) result {
		cost := present[u]      // 原价
		dCost := present[u] / 2 // 半价

		// dp0[b]: 父节点不购买，预算b的最大收益
		// dp1[b]: 父节点购买（可半价），预算b的最大收益
		dp0 := make([]int, budget+1)
		dp1 := make([]int, budget+1)

		// 子节点收益（分组背包）
		// subProfit0[b]: 子节点无折扣，预算b的最大收益
		// subProfit1[b]: 子节点有折扣（当前节点买了），预算b的最大收益
		subProfit0 := make([]int, budget+1)
		subProfit1 := make([]int, budget+1)

		uSize := cost // 当前子树最大开销（至少是当前节点的原价）

		// 分组背包：合并所有子节点
		for _, v := range g[u] {
			childResult := dfs(v)
			uSize += childResult.size // 累加子树开销

			// 倒序遍历预算，避免重复使用
			for i := budget; i >= 0; i-- {
				// 关键剪枝：只遍历到 min(childResult.size, i)
				for sub := 0; sub <= min(childResult.size, i); sub++ {
					// 子节点无折扣
					subProfit0[i] = max(subProfit0[i], subProfit0[i-sub]+childResult.dp0[sub])
					// 子节点有折扣（当前节点购买）
					subProfit1[i] = max(subProfit1[i], subProfit1[i-sub]+childResult.dp1[sub])
				}
			}
		}

		// 计算当前节点的 dp0 和 dp1
		for i := 0; i <= budget; i++ {
			// 父节点不购买的情况
			dp0[i] = subProfit0[i] // 不买当前节点
			if i >= cost {
				// 买当前节点（原价）+ 子节点享受折扣
				dp0[i] = max(dp0[i], subProfit1[i-cost]+future[u]-cost)
			}

			// 父节点购买的情况（当前节点可半价）
			dp1[i] = subProfit0[i] // 不买当前节点
			if i >= dCost {
				// 买当前节点（半价）+ 子节点享受折扣
				dp1[i] = max(dp1[i], subProfit1[i-dCost]+future[u]-dCost)
			}
		}

		return result{dp0, dp1, uSize}
	}

	return dfs(0).dp0[budget]
}
