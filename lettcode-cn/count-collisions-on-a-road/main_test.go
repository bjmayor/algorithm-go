package main

import "testing"

func TestCountCollisions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"RL碰撞", "RL", 2},
		{"RS碰撞", "RS", 1},
		{"SL碰撞", "SL", 1},
		{"RRL", "RRL", 3},
		{"RRLL", "RRLL", 4},
		{"SSSS", "SSSS", 0},
		{"RRRR", "RRRR", 0},
		{"LLLL", "LLLL", 0},
		{"R", "R", 0},
		{"L", "L", 0},
		{"S", "S", 0},
		{"", "", 0},
		{"RLRSLL", "RLRSLL", 5},
		{"LLRR", "LLRR", 0},
		{"RS", "RS", 1},
		{"SR", "SR", 0},
		{"SL", "SL", 1},
		{"LS", "LS", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countCollisions(tt.input)
			if result != tt.expected {
				t.Errorf("countCollisions(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
