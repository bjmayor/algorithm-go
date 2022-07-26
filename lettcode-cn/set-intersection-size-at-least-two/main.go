package main

import "sort"

func intersectionSizeTwo(intervals [][]int) (ans int) {
	// 先排序。按左边界升序，有边界降序
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	n, m := len(intervals), 2
	vals := make([][]int, n) // 记录集合  和 第i个 区间 已经产生交集的 数
	for i := n - 1; i >= 0; i-- {
		// 从后往前遍历。
		for j, k := intervals[i][0], len(vals[i]); k < m; k++ {
			ans++ // 区间i 的交集数还不够。得加一个
			// j是要尝试加入的数。 如果前一个区间的有边界>=j, 那么j可以加入到集合里。
			for p := i - 1; p >= 0 && intervals[p][1] >= j; p-- {
				vals[p] = append(vals[p], j)
			}
			j++
		}
	}
	return
}

func main()  {
	//println(intersectionSizeTwo([][]int{{1,2},{2,3},{2,4},{4,5} }))
	println(intersectionSizeTwo([][]int{{1,3},{1,4},{2,5},{3,5} }))
}