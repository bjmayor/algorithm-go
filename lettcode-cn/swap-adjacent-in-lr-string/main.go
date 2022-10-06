package main

import (
	"fmt"
)

//L可以向左移，R可以向右移
func canTransform(start string, end string) bool {
	var i, j int
	for i < len(start) && j < len(end) {
		for i < len(start) && start[i] == 'X' {
			i++
		}
		for j < len(end) && end[j] == 'X' {
			j++
		}
		if (i == len(start) && j < len(end)) || (i < len(start) && j == len(end)) {
			return false
		}
		if i < len(start) && j < len(end) {
			if start[i] != end[j] || (start[i] == 'L' && j > i) || (start[i] == 'R' && j < i) {
				return false
			}
			i++
			j++
		}
	}
	for i < len(start) {
		if start[i] != 'X' {
			return false
		}
		i++
	}

	for j < len(end) {
		if end[j] != 'X' {
			return false
		}
		j++
	}
	return true
}

func main() {
	tests := []struct {
		start    string
		end      string
		Expected bool
	}{
		{"XRL", "RLX", false},
		{"RXR", "XXR", false},
		{"RXXLRXRXL", "XRLXXRRLX", true},
		{"RL", "LR", false},
		{"XXXXXLXXXX", "LXXXXXXXXX", true},
		{"XXXXXLXXXLXXXX", "XXLXXXXXXXXLXX", false},
	}
	for i, t := range tests {
		var real = canTransform(t.start, t.end)
		if real != t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real)
		}
	}
}
