package two_best_non_overlapping_events

import "testing"

func TestMaxTwoEvents(t *testing.T) {
	tests := []struct {
		name   string
		events [][]int
		want   int
	}{
		{
			name:   "example 1: basic case with two non-overlapping events",
			events: [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}},
			want:   4,
		},
		{
			name:   "example 2: overlapping events",
			events: [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}},
			want:   5,
		},
		{
			name:   "example 3: three events with best combination",
			events: [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}},
			want:   8,
		},
		{
			name:   "single event",
			events: [][]int{{1, 2, 4}},
			want:   4,
		},
		{
			name:   "two completely non-overlapping events",
			events: [][]int{{1, 2, 4}, {3, 4, 3}},
			want:   7,
		},
		{
			name:   "two completely overlapping events",
			events: [][]int{{1, 5, 4}, {1, 5, 3}},
			want:   4,
		},
		{
			name:   "events with same end time",
			events: [][]int{{1, 3, 5}, {2, 3, 4}},
			want:   5,
		},
		{
			name:   "events adjacent (end time equals start time - treated as overlapping)",
			events: [][]int{{1, 3, 2}, {3, 5, 3}},
			want:   3,
		},
		{
			name:   "multiple events with best two at the end",
			events: [][]int{{1, 2, 1}, {3, 4, 2}, {5, 6, 3}, {7, 8, 4}},
			want:   7, // [5,6,3] + [7,8,4] = 7
		},
		{
			name:   "events with large values",
			events: [][]int{{1, 10, 100}, {11, 20, 200}},
			want:   300,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTwoEvents(tt.events)
			if got != tt.want {
				t.Errorf("maxTwoEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
