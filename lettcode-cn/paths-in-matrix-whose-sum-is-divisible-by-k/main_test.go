package main

import (
	"testing"
)

func TestNumberOfPaths(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}},
			k:    3,
			want: 2,
		},
		{
			name: "Example 2",
			grid: [][]int{{0, 0}},
			k:    5,
			want: 1,
		},
		{
			name: "Example 3",
			grid: [][]int{{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0}},
			k:    1,
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPaths(tt.grid, tt.k); got != tt.want {
				t.Errorf("numberOfPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
