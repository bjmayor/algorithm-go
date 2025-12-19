package main

import "sort"

// UnionFind 并查集数据结构
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind 创建并查集
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{parent: parent, rank: rank}
}

// Find 查找根节点（带路径压缩）
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

// Union 合并两个集合（按秩合并）
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return
	}

	// 按秩合并
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

// IsConnected 判断两个节点是否连通
func (uf *UnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Reset 重置节点为独立集合
func (uf *UnionFind) Reset(x int) {
	uf.parent[x] = x
	uf.rank[x] = 1
}

// findAllPeople 找出所有知晓秘密的专家
func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	// 初始化并查集
	uf := NewUnionFind(n)

	// 0 和 firstPerson 初始知道秘密
	uf.Union(0, firstPerson)

	// 按时间排序会议
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	// 按时间分组处理会议
	i := 0
	for i < len(meetings) {
		currentTime := meetings[i][2]
		peopleInMeetings := make(map[int]bool) // 当前时间参会的所有人

		// 处理同一时间的所有会议
		j := i
		for j < len(meetings) && meetings[j][2] == currentTime {
			x, y := meetings[j][0], meetings[j][1]
			uf.Union(x, y)
			peopleInMeetings[x] = true
			peopleInMeetings[y] = true
			j++
		}

		// 检查参会者是否与 0 连通，不连通的需要重置
		for person := range peopleInMeetings {
			if !uf.IsConnected(person, 0) {
				uf.Reset(person)
			}
		}

		i = j
	}

	// 收集所有与 0 连通的专家
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if uf.IsConnected(i, 0) {
			result = append(result, i)
		}
	}

	return result
}
