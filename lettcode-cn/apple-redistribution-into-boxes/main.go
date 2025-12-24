package appleredistributionintoboxes

import "slices"

func minimumBoxes(apple []int, capacity []int) int {
	totalApples := 0
	for _, a := range apple {
		totalApples += a
	}
	slices.Sort(capacity)
	slices.Reverse(capacity)
	result := 0
	for _, c := range capacity {
		result++
		totalApples -= c
		if totalApples <= 0 {
			break
		}

	}
	return result
}
