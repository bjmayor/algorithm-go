package main

import "testing"

func TestMaxRunTime(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		batteries []int
		want      int64
	}{
		{
			name:      "示例1：三块电池供两台电脑",
			n:         2,
			batteries: []int{3, 3, 3},
			want:      4,
		},
		{
			name:      "示例2：四块小电池供两台电脑",
			n:         2,
			batteries: []int{1, 1, 1, 1},
			want:      2,
		},
		{
			name:      "边界：单台电脑",
			n:         1,
			batteries: []int{10},
			want:      10,
		},
		{
			name:      "边界：电池数等于电脑数",
			n:         3,
			batteries: []int{5, 5, 5},
			want:      5,
		},
		{
			name:      "大电池被限制",
			n:         2,
			batteries: []int{10, 10, 3, 5},
			want:      14,
		},
		{
			name:      "电池全相同",
			n:         3,
			batteries: []int{4, 4, 4, 4, 4},
			want:      6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxRunTime(tt.n, tt.batteries)
			if got != tt.want {
				t.Errorf("maxRunTime(%d, %v) = %d, want %d", tt.n, tt.batteries, got, tt.want)
			}
		})
	}
}
