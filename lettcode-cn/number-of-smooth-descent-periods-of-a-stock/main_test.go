package main

import "testing"

func TestGetDescentPeriods(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int64
	}{
		{"example1", []int{3, 2, 1, 4}, 7},
		{"example2", []int{8, 6, 7, 7}, 4},
		{"example3", []int{1}, 1},
		{"all equal", []int{5, 5, 5, 5}, 4},
		{"long descent", []int{5, 4, 3, 2, 1}, 15},
		{"mixed", []int{5, 3, 2, 1, 7, 6}, 10},
		{"empty", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDescentPeriods(tt.prices)
			if got != tt.want {
				t.Fatalf("%s: got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
