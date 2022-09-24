package main

func canFormArray(arr []int, pieces [][]int) bool {
	m := map[int][]int{}
	for _, piece := range pieces {
		m[piece[0]] = piece
	}
	for i := 0; i < len(arr); {
		v := arr[i]
		if piece, ok := m[v]; ok {
			for j := 0; j < len(piece); j++ {
				if arr[i+j] != piece[j] {
					return false
				}
			}
			i += len(piece)
		} else {
			return false
		}
	}
	return true
}
