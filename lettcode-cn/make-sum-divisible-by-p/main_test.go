package main

import "testing"

func TestMinSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		p    int
		want int
	}{
		{"example1", []int{3, 1, 4, 2}, 6, 1},
		{"example2", []int{6, 3, 5, 2}, 9, 2},
		{"example3", []int{1, 2, 3}, 3, 0},
		{"single_impossible", []int{1}, 2, -1},
		{"entire_array_candidate", []int{1, 2}, 4, -1},
		{"all_same", []int{1, 1, 1}, 3, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := minSubarray(tc.nums, tc.p)
			if got != tc.want {
				t.Fatalf("%s: expected %d, got %d", tc.name, tc.want, got)
			}
		})
	}
}
