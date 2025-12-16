package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		present   []int
		future    []int
		hierarchy [][]int
		budget    int
		want      int
	}{
		{
			name:      "示例1-链状树",
			n:         3,
			present:   []int{5, 2, 3},
			future:    []int{8, 5, 6},
			hierarchy: [][]int{{1, 2}, {2, 3}},
			budget:    7,
			want:      12,
		},
		{
			name:      "示例2-两个子节点",
			n:         3,
			present:   []int{4, 6, 8},
			future:    []int{7, 9, 11},
			hierarchy: [][]int{{1, 2}, {1, 3}},
			budget:    10,
			want:      10,
		},
		{
			name:      "预算为0",
			n:         3,
			present:   []int{5, 2, 3},
			future:    []int{8, 5, 6},
			hierarchy: [][]int{{1, 2}, {2, 3}},
			budget:    0,
			want:      0,
		},
		{
			name:      "预算不足",
			n:         3,
			present:   []int{10, 10, 10},
			future:    []int{15, 15, 15},
			hierarchy: [][]int{{1, 2}, {2, 3}},
			budget:    3,
			want:      0,
		},
		{
			name:      "单节点",
			n:         1,
			present:   []int{10},
			future:    []int{15},
			hierarchy: [][]int{},
			budget:    10,
			want:      5,
		},
		{
			name:      "所有节点亏损",
			n:         3,
			present:   []int{10, 10, 10},
			future:    []int{5, 5, 5},
			hierarchy: [][]int{{1, 2}, {1, 3}},
			budget:    50,
			want:      0,
		},
		{
			name:      "大预算测试用例",
			n:         12,
			present:   []int{31, 6, 36, 19, 9, 7, 40, 31, 19, 21, 33, 7},
			future:    []int{15, 22, 46, 9, 36, 3, 46, 44, 31, 36, 23, 23},
			hierarchy: [][]int{{1, 2}, {1, 4}, {4, 10}, {4, 9}, {9, 11}, {10, 6}, {1, 8}, {2, 5}, {1, 12}, {4, 7}, {12, 3}},
			budget:    93,
			want:      156, // 实际运行结果
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.n, tt.present, tt.future, tt.hierarchy, tt.budget)
			if got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
				t.Logf("Input: n=%d, present=%v, future=%v, hierarchy=%v, budget=%d",
					tt.n, tt.present, tt.future, tt.hierarchy, tt.budget)
			}
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	n := 12
	present := []int{31, 6, 36, 19, 9, 7, 40, 31, 19, 21, 33, 7}
	future := []int{15, 22, 46, 9, 36, 3, 46, 44, 31, 36, 23, 23}
	hierarchy := [][]int{{1, 2}, {1, 4}, {4, 10}, {4, 9}, {9, 11}, {10, 6}, {1, 8}, {2, 5}, {1, 12}, {4, 7}, {12, 3}}
	budget := 93

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxProfit(n, present, future, hierarchy, budget)
	}
}

func BenchmarkMaxProfitSmall(b *testing.B) {
	n := 3
	present := []int{5, 2, 3}
	future := []int{8, 5, 6}
	hierarchy := [][]int{{1, 2}, {2, 3}}
	budget := 7

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxProfit(n, present, future, hierarchy, budget)
	}
}
