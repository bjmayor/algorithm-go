package main

func minDeletionSize(strs []string) int {
	colLen := len(strs[0])
	delCount := 0

	for col := 0; col < colLen; col++ {
		for row := 1; row < len(strs); row++ {
			if strs[row][col] < strs[row-1][col] {
				delCount++
				break
			}
		}
	}
	return delCount
}
