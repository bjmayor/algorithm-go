package main

func minimumOperations(nums []int) int {
	result := 0
	for _, v := range nums {
		if v%3 != 0 {
			result++
		}
	}
	return result
}

func main() {
	// Test cases using table-driven approach
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"example1", []int{1, 2, 3, 4, 5, 6}, 4},
		{"all_divisible", []int{3, 6, 9, 12}, 0},
		{"none_divisible", []int{1, 2, 4, 5, 7, 8}, 6},
		{"empty", []int{}, 0},
		{"single_divisible", []int{9}, 0},
		{"single_not_divisible", []int{5}, 1},
		{"mixed_large", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 8},
	}

	allPassed := true
	for _, test := range tests {
		result := minimumOperations(test.input)
		if result != test.expected {
			println("FAIL:", test.name, "got", result, "expected", test.expected)
			allPassed = false
		}
	}

	if allPassed {
		println("all passed")
	}
}
