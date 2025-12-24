package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		strategy []int
		k        int
		expected int64
	}{
		// Example 1 from the problem description
		{
			name:     "Example 1 - Modify first two elements",
			prices:   []int{4, 2, 8},
			strategy: []int{-1, 0, 1},
			k:        2,
			expected: 10,
		},
		// Example 2 from the problem description
		{
			name:     "Example 2 - No modification needed",
			prices:   []int{5, 4, 3},
			strategy: []int{1, 1, 0},
			k:        2,
			expected: 9,
		},
		// Edge cases
		{
			name:     "Minimum length with k=2",
			prices:   []int{1, 2},
			strategy: []int{-1, 1},
			k:        2,
			expected: 2, // Original: (-1*1)+(1*2)=1, Modified: (0*1)+(1*2)=2
		},
		{
			name:     "All negative strategies",
			prices:   []int{3, 4, 5, 6},
			strategy: []int{-1, -1, -1, -1},
			k:        2,
			expected: -1, // Best to change to [0,1,-1,-1] or [-1,0,1,-1]
		},
		{
			name:     "All positive strategies",
			prices:   []int{3, 4, 5, 6},
			strategy: []int{1, 1, 1, 1},
			k:        2,
			expected: 18, // Keep as is since all positive
		},
		{
			name:     "All zero strategies",
			prices:   []int{3, 4, 5, 6},
			strategy: []int{0, 0, 0, 0},
			k:        2,
			expected: 6, // Best to change to [0,1,0,0] or [0,0,1,0]
		},
		// Mixed strategies
		{
			name:     "Mixed strategies - best modification in middle",
			prices:   []int{10, 1, 10, 1, 10},
			strategy: []int{-1, 0, -1, 0, 1},
			k:        4,
			expected: 21, // Change middle 4 elements to maximize profit
		},
		{
			name:     "Large prices with small k",
			prices:   []int{100, 1, 100, 1, 100},
			strategy: []int{-1, 1, -1, 1, -1},
			k:        2,
			expected: -99, // Find best 2-element segment to modify
		},
		// No benefit from modification
		{
			name:     "Already optimal - modification decreases profit",
			prices:   []int{5, 4, 3, 2, 1},
			strategy: []int{1, 1, 1, 1, 1},
			k:        4,
			expected: 15, // Keep original strategy
		},
		// Equal k to array length
		{
			name:     "K equals array length",
			prices:   []int{1, 2, 3, 4},
			strategy: []int{-1, 0, 0, 1},
			k:        4,
			expected: 7, // Change to [0,0,1,1]
		},
		// Large array test
		{
			name:     "Large array with alternating prices",
			prices:   []int{1, 10, 1, 10, 1, 10, 1, 10},
			strategy: []int{-1, 1, -1, 1, -1, 1, -1, 1},
			k:        6,
			expected: 36, // Optimal modification of 6 elements
		},
		// Zero and negative price scenarios (if allowed)
		{
			name:     "With price 1 and strategy mix",
			prices:   []int{1, 1, 1, 1},
			strategy: []int{-1, 0, 1, -1},
			k:        2,
			expected: 1, // Best modification of 2 elements
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices, tt.strategy, tt.k)
			if got != tt.expected {
				t.Errorf("maxProfit(%v, %v, %d) = %d, want %d",
					tt.prices, tt.strategy, tt.k, got, tt.expected)
			}
		})
	}
}
