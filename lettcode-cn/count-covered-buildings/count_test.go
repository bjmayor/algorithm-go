package main

import (
	"testing"
)

func TestCountCoveredBuildings(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		buildings [][]int
		expected  int
	}{
		{
			name: "十字形建筑",
			n:    5,
			buildings: [][]int{
				{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4},
				{0, 2}, {1, 2}, {3, 2}, {4, 2},
			},
			expected: 1,
		},
		{
			name: "单个建筑",
			n:    1,
			buildings: [][]int{
				{0, 0},
			},
			expected: 0,
		},
		{
			name: "一行建筑",
			n:    5,
			buildings: [][]int{
				{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4},
			},
			expected: 0, // 一行建筑没有上下方向，不被覆盖
		},
		{
			name: "一列建筑",
			n:    5,
			buildings: [][]int{
				{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
			},
			expected: 0, // 一列建筑没有左右方向，不被覆盖
		},
		{
			name: "3x3网格",
			n:    3,
			buildings: [][]int{
				{0, 0}, {0, 1}, {0, 2},
				{1, 0}, {1, 1}, {1, 2},
				{2, 0}, {2, 1}, {2, 2},
			},
			expected: 1, // 只有中心点 (1,1) 被覆盖
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countCoveredBuildings(tt.n, tt.buildings)
			if result != tt.expected {
				t.Errorf("countCoveredBuildings(%d, %v) = %d, expected %d", tt.n, tt.buildings, result, tt.expected)
			}
		})
	}
}

// 基准测试
func BenchmarkCountCoveredBuildings(b *testing.B) {
	buildings := [][]int{
		{0, 0}, {0, 1}, {0, 2},
		{1, 0}, {1, 1}, {1, 2},
		{2, 0}, {2, 1}, {2, 2},
	}
	n := 3

	for i := 0; i < b.N; i++ {
		countCoveredBuildings(n, buildings)
	}
}