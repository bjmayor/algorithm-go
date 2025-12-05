package main

import "testing"

func TestCountNumberOfTrapezoids(t *testing.T) {
	tests := []struct {
		name string
		pts  [][]int
		want int
	}{
		{"example1", [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}}, 2},
		{"example2", [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}}, 1},
		{"parallelogram", [][]int{{0, 0}, {2, 0}, {2, 1}, {0, 1}}, 1},
		{"no trapezoid - collinear 4", [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, 0},
		{"more points", [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}, {3, 0}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTrapezoids(tt.pts)
			if got != tt.want {
				t.Fatalf("%s: got %d want %d", tt.name, got, tt.want)
			}
		})
	}
}
