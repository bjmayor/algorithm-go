package main

/*
n == strs.length
1 <= n <= 100
1 <= strs[i].length <= 100
strs[i] 由小写英文字母组成
*/
func minDeletionSize(strs []string) int {
	n := len(strs)
	m := len(strs[0])
	ans := 0
	sorted := make([]bool, n-1) // sorted[i] true => strs[i] already <= strs[i+1] due to previous columns
	for col := 0; col < m; col++ {
		deleteCol := false
		for row := 0; row < n-1; row++ {
			if sorted[row] {
				continue
			}
			if strs[row][col] > strs[row+1][col] {
				ans++
				deleteCol = true
				break
			}
		}
		if deleteCol {
			continue
		}
		for row := 0; row < n-1; row++ {
			if !sorted[row] && strs[row][col] < strs[row+1][col] {
				sorted[row] = true
			}
		}
	}
	return ans
}
