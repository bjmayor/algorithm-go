package main

func maxChunksToSorted(arr []int) int {
	cur := -1
	ans := 0
	for i, v := range arr {
		cur = max(cur, v)
		if cur <= i {
			ans++
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	println(maxChunksToSorted([]int{4, 3, 2, 1, 0}))
	println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
}
