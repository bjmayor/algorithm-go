package two_best_non_overlapping_events

import "sort"

func maxTwoEvents(events [][]int) int {
	result := 0
	maxValue := make([]int, len(events)) // maxValue[i] 表示前i个事件中，单个事件的最大价值
	// 按照活动结束时间排序
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	maxValue[0] = events[0][2]
	result = events[0][2]
	for i := 1; i < len(events); i++ {
		currentValue := events[i][2]
		// 二分查找第一个结束时间小于当前活动开始时间的活动
		left, right := 0, i-1
		for left <= right {
			mid := (right + left) / 2
			// mid 满足条件
			if events[mid][1] < events[i][0] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if right >= 0 {
			// 找到不重叠的事件，加上前面单个事件的最大价值
			currentValue += maxValue[right]
		}
		// 更新前i个事件中单个事件的最大价值
		maxValue[i] = max(maxValue[i-1], events[i][2])
		result = max(result, currentValue)
	}
	return result
}
