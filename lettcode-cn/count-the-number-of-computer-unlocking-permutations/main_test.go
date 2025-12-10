package main

import "testing"

func TestCountPermutations(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"example1", []int{1, 2, 3}, 2},
		{"example2", []int{3, 3, 3, 4, 4, 4}, 0},
		{"eq_nonsolving", []int{1, 1, 2}, 0},
		{"duplicated_but_valid", []int{1, 2, 2}, 2},
		{"n2_valid", []int{1, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPermutations(tt.input)
			if got != tt.want {
				t.Fatalf("countPermutations(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
