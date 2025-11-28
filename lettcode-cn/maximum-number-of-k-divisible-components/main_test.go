package main

import "testing"

func TestMaxKDivisibleComponents(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		edges  [][]int
		values []int
		k      int
		want   int
	}{
		// 官方示例 1（来自题目附件）
		{
			name: "官方示例1",
			n:    5,
			edges: [][]int{
				{0, 2}, {1, 2}, {1, 3}, {2, 4},
			},
			values: []int{1, 8, 1, 4, 4},
			k:      6,
			want:   2,
		},
		// 官方示例 2（来自题目附件）
		{
			name: "官方示例2",
			n:    7,
			edges: [][]int{
				{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6},
			},
			values: []int{3, 0, 6, 1, 5, 2, 1},
			k:      3,
			want:   3,
		},
		// 额外示例：验证贪心策略
		{
			name: "贪心示例: 基本情况",
			n:    6,
			edges: [][]int{
				{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5},
			},
			values: []int{1, 2, 3, 4, 5, 6},
			k:      3,
			want:   4, // 贪心结果：{2},{5},{3,4},{0,1}
		},
		{
			name:   "示例2: 单节点能整除",
			n:      1,
			edges:  [][]int{},
			values: []int{5},
			k:      5,
			want:   1,
		},
		{
			name:   "示例3: 单节点不能整除",
			n:      1,
			edges:  [][]int{},
			values: []int{7},
			k:      5,
			want:   0, // 没有 k-可整除块
		},
		{
			name: "示例4: 所有节点单独成块",
			n:    4,
			edges: [][]int{
				{0, 1}, {1, 2}, {2, 3},
			},
			values: []int{3, 6, 9, 12},
			k:      3,
			want:   4, // 每个节点都能被3整除
		},
		{
			name: "示例5: 整棵树是一个块",
			n:    3,
			edges: [][]int{
				{0, 1}, {0, 2},
			},
			values: []int{1, 2, 3},
			k:      5,
			want:   0, // 总和=6, 无 k-可整除块
		},
		{
			name: "示例6: 链式树结构",
			n:    5,
			edges: [][]int{
				{0, 1}, {1, 2}, {2, 3}, {3, 4},
			},
			values: []int{2, 1, 2, 5, 5},
			k:      5,
			want:   3, // {4}=5, {3}=5, {0,1,2}=5
		},
		{
			name: "示例7: 负数值",
			n:    4,
			edges: [][]int{
				{0, 1}, {1, 2}, {1, 3},
			},
			values: []int{3, 0, 3, 6},
			k:      3,
			want:   4, // 每个节点/子树都能被3整除
		},
		{
			name: "示例8: 星型树",
			n:    5,
			edges: [][]int{
				{0, 1}, {0, 2}, {0, 3}, {0, 4},
			},
			values: []int{2, 3, 5, 5, 5},
			k:      5,
			want:   4, // {2},{3},{4}单点可整除且{0,1}=2+3=5，共4个块
		},
		{
			name: "示例9: 复杂树结构",
			n:    7,
			edges: [][]int{
				{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6},
			},
			values: []int{3, 0, 6, 1, 5, 2, 1},
			k:      3,
			want:   3, // {5,6,2}=9, {3,4,1}=6, {0}=3
		},
		{
			name: "示例10: k=1时所有和都可整除",
			n:    3,
			edges: [][]int{
				{0, 1}, {1, 2},
			},
			values: []int{5, 7, 3},
			k:      1,
			want:   3, // 所有数都能被1整除
		},
		{
			name: "示例11: 大k值无法分割",
			n:    4,
			edges: [][]int{
				{0, 1}, {1, 2}, {2, 3},
			},
			values: []int{1, 1, 1, 1},
			k:      100,
			want:   0, // 总和=4, 无 k-可整除块
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxKDivisibleComponents(tt.n, tt.edges, tt.values, tt.k)
			if got != tt.want {
				t.Errorf("maxKDivisibleComponents() = %v, want %v", got, tt.want)
				t.Logf("Input: n=%d, edges=%v, values=%v, k=%d", tt.n, tt.edges, tt.values, tt.k)
			}
		})
	}
}

// 基准测试
func BenchmarkMaxKDivisibleComponents(b *testing.B) {
	// 构造一个较大的测试用例
	n := 100
	edges := make([][]int, 0, n-1)
	values := make([]int, n)

	// 构造链式结构
	for i := 0; i < n-1; i++ {
		edges = append(edges, []int{i, i + 1})
		values[i] = i % 10
	}
	values[n-1] = (n - 1) % 10
	k := 3

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxKDivisibleComponents(n, edges, values, k)
	}
}
