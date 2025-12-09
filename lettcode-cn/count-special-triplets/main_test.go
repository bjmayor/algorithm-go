package main

import "testing"

func TestSpecialTriplets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "示例1",
			nums: []int{6, 3, 6},
			want: 1,
		},
		{
			name: "示例2",
			nums: []int{0, 1, 0, 0},
			want: 1,
		},
		{
			name: "示例3",
			nums: []int{8, 4, 2, 8, 4},
			want: 2,
		},
		{
			name: "最小用例",
			nums: []int{4, 2, 4},
			want: 1,
		},
		{
			name: "无特殊三元组",
			nums: []int{1, 2, 3},
			want: 0,
		},
		{
			name: "多个相同值",
			nums: []int{2, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "包含0的情况",
			nums: []int{0, 0, 0, 0},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := specialTriplets(tt.nums)
			if got != tt.want {
				t.Errorf("specialTriplets(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
