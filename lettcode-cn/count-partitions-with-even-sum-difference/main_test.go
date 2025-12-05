package main

import "testing"

func TestCountPartitionsDiff(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"示例1", []int{10, 10, 3, 7, 6}, 4},
		{"示例2", []int{1, 2, 2}, 0},
		{"示例3", []int{2, 4, 6, 8}, 3},
		{"全偶数", []int{2, 4, 6}, 2},
		{"全奇数差值偶数", []int{1, 3, 5, 7}, 3},
		{"两元素差值偶数", []int{1, 3}, 1},
		{"两元素差值奇数", []int{1, 2}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countPartitionsDiff(tt.nums)
			if result != tt.expected {
				t.Errorf("countPartitionsDiff(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
