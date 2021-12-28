package main

import "math"

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	// 先记录下有哪些停靠站点 有公交车
	// 转换成 公交车为节点。。这样就变成了多起点/多终点的图问题。 2辆车有共同的节点时，表示有变，权值为1
	m := make(map[int][]int)
	n := len(routes)
	edge := make([][]int, n) // 存下公交车能连到的边
	for i := range edge {
		edge[i] = make([]int, 0)
	}
	for bus, stations:= range routes {
		for _, station:= range stations {

			for _, j := range m[station] {
				edge[bus] = append(edge[bus], j)
				edge[j] = append(edge[j], bus)
			}
			m[station] = append(m[station], bus)
		}
	}
	dis := make([]int, n)// 因为是多起点和多终点，所有会有多个解，需要选择最小的
	for i := range dis {
		dis[i] = -1
	}
	q := []int{}
	for _, bus := range m[source] {
		q =  append(q, bus)
		dis[bus] = 1
	}
	for len(q) > 0 {
		h := q[0]
		q = q[1:]
		// 从h 换到 bus 上。。b==true 表示可以换乘
		for _, bus := range edge[h] {
			if  dis[bus] == -1 {
				dis[bus] = dis[h] + 1
				q = append(q, bus)
			}
		}
	}
	ans := math.MaxInt32
	// 每个终点 检查是否可到达，并选取最先的那个
	for _, bus := range m[target] {
		if dis[bus] != -1 && dis[bus] < ans {
			ans = dis[bus]
		}
	}
	if ans < math.MaxInt32 {
		return ans
	}
	return -1

}

func main() {
	println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))                                                                                                                                                                           // 2
	println(numBusesToDestination([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12))                                                                                                                                              // -1
	println(numBusesToDestination([][]int{{1, 9, 12, 20, 23, 24, 35, 38}, {10, 21, 24, 31, 32, 34, 37, 38, 43}, {10, 19, 28, 37}, {8}, {14, 19}, {11, 17, 23, 31, 41, 43, 44}, {21, 26, 29, 33}, {5, 11, 33, 41}, {4, 5, 8, 9, 24, 44}}, 37, 28)) // 1
}
