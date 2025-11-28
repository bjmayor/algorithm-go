package main

func findFinalValue(nums []int, original int) int {
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	for {
		if _, exists := numSet[original]; exists {
			original *= 2
		} else {
			break
		}
	}

	return original
}
