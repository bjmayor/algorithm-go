package main

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			// 右侧
			// 之前的交集插入进去
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			// 左侧无交集，
			ans = append(ans, interval)
		} else {
			// 有交集，更新
			left = min(left, interval[0])
			right = max(right, interval[1])
		}

	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
